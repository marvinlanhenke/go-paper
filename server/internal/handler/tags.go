package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/marvinlanhenke/go-paper/internal/repository"
	"github.com/marvinlanhenke/go-paper/internal/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TagPayload struct {
	Name string `json:"name" validate:"required,max=100"`
}

type tagHandler struct {
	logger     *zap.SugaredLogger
	repository *repository.Repository
}

func NewTagHandler(logger *zap.SugaredLogger, repository *repository.Repository) *tagHandler {
	return &tagHandler{logger: logger, repository: repository}
}

func (h *tagHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload TagPayload

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

	tag := &repository.Tag{Name: payload.Name}

	if err := h.repository.Tags.Create(r.Context(), tag); err != nil {
		h.logger.Errorw("failed to create tag", "error", err)
		utils.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, tag)
}

func (h *tagHandler) Read(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.logger.Errorw("failed to convert id param", "error", err)
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tag, err := h.repository.Tags.Read(r.Context(), id)
	if err != nil {
		h.logger.Errorw("failed to read tag", "error", err)

		if err == gorm.ErrRecordNotFound {
			utils.JSONError(w, http.StatusNotFound, err.Error())
			return
		}

		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, tag)
}

func (h *tagHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.logger.Errorw("failed to convert id param", "error", err)
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.repository.Tags.Delete(r.Context(), id); err != nil {
		h.logger.Errorw("failed to delete tag", "error", err)

		if err == gorm.ErrRecordNotFound {
			utils.JSONError(w, http.StatusNotFound, err.Error())
			return
		}

		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
