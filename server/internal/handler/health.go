package handler

import (
	"encoding/json"
	"net/http"

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&envelope{data}); err != nil {
		h.logger.Errorw("error while encoding JSON message", "error", err)
	}
}
