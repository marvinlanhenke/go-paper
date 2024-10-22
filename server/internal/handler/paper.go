package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/marvinlanhenke/go-paper/internal/repository"
	"github.com/marvinlanhenke/go-paper/internal/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type paperKey string

const paperCtx paperKey = "paperCtx"

type CreatePaperPayload struct {
	Title       string `json:"title" validate:"required,max=255"`
	URL         string `json:"url" validate:"required,max=255"`
	Description string `json:"description"`
}

type UpdatePaperPayload struct {
	Title       *string `json:"title" validate:"omitempty,max=255"`
	URL         *string `json:"url" validate:"omitempty,max=255"`
	Description *string `json:"description"`
	IsRead      *bool   `json:"is_read"`
}

type paperHandler struct {
	logger     *zap.SugaredLogger
	repository *repository.Repository
}

func NewPaperHandler(logger *zap.SugaredLogger, repository *repository.Repository) *paperHandler {
	return &paperHandler{logger: logger, repository: repository}
}

func (h *paperHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload CreatePaperPayload

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		h.logger.Errorw("failed to decode body", "error", err)
		utils.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		h.logger.Errorw("failed to validate payload", "error", err)
		utils.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	paper := &repository.Paper{Title: payload.Title, Description: payload.Description, URL: payload.URL}

	if err := h.repository.Papers.Create(r.Context(), paper); err != nil {
		h.logger.Errorw("failed to create tag", "error", err)
		utils.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, paper)
}

func (h *paperHandler) Read(w http.ResponseWriter, r *http.Request) {
	paper := getPaperFromCtx(r)
	utils.JSONResponse(w, http.StatusOK, paper)
}

func (h *paperHandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	papers, err := h.repository.Papers.ReadAll(r.Context())
	if err != nil {
		h.logger.Errorw("failed to read all entities", "error", err)
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, papers)
}

func (h *paperHandler) Update(w http.ResponseWriter, r *http.Request) {
	paper := getPaperFromCtx(r)

	var payload UpdatePaperPayload

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		h.logger.Errorw("failed to decode body", "error", err)
		utils.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		h.logger.Errorw("failed to validate payload", "error", err)
		utils.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if payload.Title != nil {
		paper.Title = *payload.Title
	}
	if payload.Description != nil {
		paper.Description = *payload.Description
	}
	if payload.URL != nil {
		paper.URL = *payload.URL
	}
	if payload.IsRead != nil {
		paper.IsRead = *payload.IsRead
	}

	if err := h.repository.Papers.Update(r.Context(), paper); err != nil {
		h.logger.Errorw("failed to update entity", "error", err)
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, paper)
}

func (h *paperHandler) Delete(w http.ResponseWriter, r *http.Request) {
	paper := getPaperFromCtx(r)

	if err := h.repository.Papers.Delete(r.Context(), paper); err != nil {
		h.logger.Errorw("failed to delete entity", "error", err)
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, nil)
}

func (h *paperHandler) WithPaperContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			h.logger.Errorw("failed to convert id param", "error", err)
			utils.JSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		paper, err := h.repository.Papers.Read(r.Context(), id)
		if err != nil {
			h.logger.Errorw("failed to read entity", "error", err)

			if err == gorm.ErrRecordNotFound {
				utils.JSONError(w, http.StatusNotFound, err.Error())
				return
			}

			utils.JSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), paperCtx, paper)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPaperFromCtx(r *http.Request) *repository.Paper {
	paper, _ := r.Context().Value(paperCtx).(*repository.Paper)
	return paper
}
