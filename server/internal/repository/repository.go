package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	Papers interface {
		Create(context.Context, *Paper) error
		Read(context.Context, int) (*Paper, error)
		ReadAll(context.Context) ([]Paper, error)
		Update(context.Context, *Paper) error
		Delete(context.Context, *Paper) error
	}
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		Papers: &paperRepository{DB: db},
	}
}
