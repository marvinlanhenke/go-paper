package handler

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type HealthCheckHandler struct {
	Logger *zap.SugaredLogger
}

func NewHealthCheckHandler(logger *zap.SugaredLogger) *HealthCheckHandler {
	return &HealthCheckHandler{Logger: logger}
}

func (h *HealthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
	}

	type envelope struct {
		Data any `json:"data"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&envelope{data}); err != nil {
		h.Logger.Errorw("error while encoding JSON message", "error", err)
	}
}
