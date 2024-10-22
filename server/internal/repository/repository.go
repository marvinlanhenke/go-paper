package repository

import (
	"context"

	"github.com/marvinlanhenke/go-paper/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	Tags interface {
		Create(context.Context, *model.Tag) error
	}
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		Tags: &tagRepository{DB: db},
	}
}
