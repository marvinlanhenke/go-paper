package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marvinlanhenke/go-paper/internal/handler"
	"github.com/marvinlanhenke/go-paper/internal/repository"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreatePaperHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "expected no error, instead got %v", err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	require.NoError(t, err, "expected no error, instead got %v", err)

	repository := repository.New(gormDB)
	logger := zap.NewNop().Sugar()
	paperHandler := handler.NewPaperHandler(logger, repository)

	payload := handler.CreatePaperPayload{
		Title:       "Test Paper",
		URL:         "http://example.com",
		Description: "Test Description",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "papers" \("title","description","url","is_read","created_at","updated_at","deleted_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6,\$7\) RETURNING "id"`).
		WithArgs(
			payload.Title,
			payload.Description,
			payload.URL,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1)) // Mocked returned ID
	mock.ExpectCommit()

	jsonPayload, err := json.Marshal(payload)
	require.NoError(t, err, "expected no error, instead got %v", err)

	req, err := http.NewRequest(http.MethodPost, "/v1/papers", bytes.NewReader(jsonPayload))
	require.NoError(t, err, "expected no error, instead got %v", err)

	rr := httptest.NewRecorder()
	paperHandler.Create(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code, "expected %v, instead got %v", http.StatusCreated, rr.Code)
	require.NoError(t, mock.ExpectationsWereMet())
}
