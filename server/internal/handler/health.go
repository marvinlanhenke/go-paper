package handler

import (
	"net/http"

	"github.com/marvinlanhenke/go-paper/internal/utils"
)

type healthCheckHandler struct {
	env     string
	version string
}

func NewHealthCheckHandler(env, version string) *healthCheckHandler {
	return &healthCheckHandler{env: env, version: version}
}

func (h *healthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "ok",
		"environment": h.env,
		"version":     h.version,
	}

	type envelope struct {
		Data any `json:"data"`
	}

	utils.JSONResponse(w, http.StatusOK, &envelope{data})
}
