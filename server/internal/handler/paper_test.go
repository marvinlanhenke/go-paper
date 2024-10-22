package handler_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi/v5"
	"github.com/marvinlanhenke/go-paper/internal/handler"
	"github.com/marvinlanhenke/go-paper/internal/repository"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createFixture(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *handler.PaperHandler) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "expected no error, instead got %v", err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	require.NoError(t, err, "expected no error, instead got %v", err)

	repository := repository.New(gormDB)
	// logger := zap.NewNop().Sugar()
	logger := zap.Must(zap.NewProduction()).Sugar()
	paperHandler := handler.NewPaperHandler(logger, repository)

	return db, mock, paperHandler
}

func TestCreatePaperHandler(t *testing.T) {
	db, mock, paperHandler := createFixture(t)
	defer db.Close()

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

func TestUpdatePaperHandler(t *testing.T) {
	db, mock, paperHandler := createFixture(t)
	defer db.Close()

	existingPaper := &repository.Paper{
		ID:          1,
		Title:       "Original Title",
		Description: "Original Description",
		URL:         "http://original.com",
		IsRead:      false,
	}

	mock.ExpectQuery(`SELECT .* FROM "papers" WHERE "papers"."id" = \$1`).
		WithArgs(existingPaper.ID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "url", "is_read", "created_at", "updated_at", "deleted_at"}).
			AddRow(existingPaper.ID, existingPaper.Title, existingPaper.Description, existingPaper.URL, existingPaper.IsRead, existingPaper.CreatedAt, existingPaper.UpdatedAt, existingPaper.DeletedAt))

	newTitle := "Updated Title"
	newDescription := "Updated Description"
	newURL := "http://updated.com"
	newIsRead := true

	payload := handler.UpdatePaperPayload{
		Title:       &newTitle,
		Description: &newDescription,
		URL:         &newURL,
		IsRead:      &newIsRead,
	}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "papers" SET "id"=\$1,"title"=\$2,"description"=\$3,"url"=\$4,"is_read"=\$5,"updated_at"=\$6 WHERE id = \$7`).
		WithArgs(
			existingPaper.ID,
			newTitle,
			newDescription,
			newURL,
			newIsRead,
			sqlmock.AnyArg(), // 'updated_at' timestamp
			existingPaper.ID,
		).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected
	mock.ExpectCommit()

	jsonPayload, err := json.Marshal(payload)
	require.NoError(t, err, "expected no error, instead got %v", err)

	req, err := http.NewRequest(http.MethodPatch, "/v1/papers/1", bytes.NewReader(jsonPayload))
	require.NoError(t, err, "expected no error, instead got %v", err)

	r := chi.NewRouter()
	r.Route("/v1/papers/{id}", func(r chi.Router) {
		r.Use(paperHandler.WithPaperContext)
		r.Patch("/", paperHandler.Update)
	})
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code, "expected %v, instead got %v", http.StatusOK, rr.Code)
	require.NoError(t, mock.ExpectationsWereMet())
}
