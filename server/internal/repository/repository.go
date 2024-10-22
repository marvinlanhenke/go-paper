package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	Tags interface {
		Create(context.Context, *Tag) error
	}
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		Tags: &tagRepository{DB: db},
	}
}
