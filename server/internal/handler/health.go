package handler

import (
	"net/http"

	"github.com/marvinlanhenke/go-paper/internal/utils"
	"go.uber.org/zap"
)

type healthCheckHandler struct {
	logger *zap.SugaredLogger
}

func NewHealthCheckHandler(logger *zap.SugaredLogger) *healthCheckHandler {
	return &healthCheckHandler{logger: logger}
}

func (h *healthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
	}

	type envelope struct {
		Data any `json:"data"`
	}

	if err := utils.WriteJSON(w, http.StatusOK, &envelope{data}); err != nil {
		h.logger.Errorw("error while writing JSON response", "error", err)
	}
}
