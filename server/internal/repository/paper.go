package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Paper struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text"`
	URL         string    `json:"url" gorm:"type:varchar(255);not null"`
	IsRead      bool      `json:"is_read" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type paperRepository struct {
	DB *gorm.DB
}

func (r *paperRepository) Create(ctx context.Context, paper *Paper) error {
	return r.DB.WithContext(ctx).Create(paper).Error
}

func (r *paperRepository) Read(ctx context.Context, paperID int) (*Paper, error) {
	var paper Paper

	if err := r.DB.WithContext(ctx).First(&paper, paperID).Error; err != nil {
		return nil, err
	}

	return &paper, nil
}

func (r *paperRepository) ReadAll(ctx context.Context) ([]Paper, error) {
	var papers []Paper

	if err := r.DB.WithContext(ctx).Find(&papers).Error; err != nil {
		return nil, err
	}

	return papers, nil
}

func (r *paperRepository) Update(ctx context.Context, paper *Paper) error {
	return r.DB.WithContext(ctx).Model(&Paper{}).Where("id = ?", paper.ID).Updates(paper).Error
}

func (r *paperRepository) Delete(ctx context.Context, paper *Paper) error {
	return r.DB.WithContext(ctx).Delete(paper).Error
}
