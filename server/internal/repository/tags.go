package repository

import (
	"context"

	"github.com/marvinlanhenke/go-paper/internal/model"
	"gorm.io/gorm"
)

type tagRepository struct {
	DB *gorm.DB
}

func (r *tagRepository) Create(ctx context.Context, tag *model.Tag) error {
	return r.DB.Create(tag).Error
}
