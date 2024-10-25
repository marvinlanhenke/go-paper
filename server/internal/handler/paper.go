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

type PaperHandler struct {
	logger     *zap.SugaredLogger
	repository *repository.Repository
}

func NewPaperHandler(logger *zap.SugaredLogger, repository *repository.Repository) *PaperHandler {
	return &PaperHandler{logger: logger, repository: repository}
}

// Create creates a new paper
//
// @Summary Create a new paper
// @Description Create a new paper with the input payload
// @Tags papers
// @Accept json
// @Produce json
// @Param paper body CreatePaperPayload true "Create Paper"
// @Success 201 {object} repository.Paper
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /papers [post]
func (h *PaperHandler) Create(w http.ResponseWriter, r *http.Request) {
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
		h.logger.Errorw("failed to create entity", "error", err)
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, paper)
}

// Read retrieves a specific paper by its ID.
//
// @Summary Get a Paper by ID
// @Description Retrieve the details of a paper using its unique ID.
// @Tags papers
// @Accept json
// @Produce json
// @Param id path int true "Paper ID"
// @Success 200 {object} repository.Paper
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /papers/{id} [get]
func (h *PaperHandler) Read(w http.ResponseWriter, r *http.Request) {
	paper := getPaperFromCtx(r)
	utils.JSONResponse(w, http.StatusOK, paper)
}

// ReadAll retrieves all papers.
//
// @Summary Get All Papers
// @Description Retrieve a list of all papers stored in the system.
// @Tags papers
// @Accept json
// @Produce json
// @Success 200 {array} repository.Paper
// @Failure 500 {object} error
// @Router /papers [get]
func (h *PaperHandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	papers, err := h.repository.Papers.ReadAll(r.Context())
	if err != nil {
		h.logger.Errorw("failed to read all entities", "error", err)
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, papers)
}

// Update modifies an existing paper by its ID.
//
// @Summary Update a Paper by ID
// @Description Update the details of a paper using its unique ID.
// @Tags papers
// @Accept json
// @Produce json
// @Param id path int true "Paper ID"
// @Param paper body handler.UpdatePaperPayload true "Update Paper"
// @Success 200 {object} repository.Paper
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /papers/{id} [patch]
func (h *PaperHandler) Update(w http.ResponseWriter, r *http.Request) {
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

func (h *PaperHandler) Delete(w http.ResponseWriter, r *http.Request) {
	paper := getPaperFromCtx(r)

	if err := h.repository.Papers.Delete(r.Context(), paper); err != nil {
		h.logger.Errorw("failed to delete entity", "error", err)
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, nil)
}

func (h *PaperHandler) WithPaperContext(next http.Handler) http.Handler {
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
