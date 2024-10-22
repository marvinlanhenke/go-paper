package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marvinlanhenke/go-paper/internal/handler"
	"github.com/stretchr/testify/require"
)

const (
	env     = "development"
	version = "0.0.1"
)

func TestHealthCheckHandler(t *testing.T) {
	handler := handler.NewHealthCheckHandler(env, version)

	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/health", nil)
	require.NoError(t, err, "expected no error, instead got %v", err)

	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusOK, rr.Code, "expected %v, instead got %v", http.StatusOK, rr.Code)

	expected := `{"data":{"status":"ok", "environment":"development", "version":"0.0.1"}}`
	result := rr.Body.String()
	require.JSONEq(t, expected, result, "expected %v, instead got %v", expected, result)
}
