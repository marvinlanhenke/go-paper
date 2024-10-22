package handler

import (
	"net/http"

	"github.com/marvinlanhenke/go-paper/internal/repository"
	"github.com/marvinlanhenke/go-paper/internal/utils"
	"go.uber.org/zap"
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
