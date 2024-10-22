package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type tagRepository struct {
	DB *gorm.DB
}

func (r *tagRepository) Create(ctx context.Context, tag *Tag) error {
	return r.DB.WithContext(ctx).Create(tag).Error
}

func (r *tagRepository) Read(ctx context.Context, tagID int) (*Tag, error) {
	var tag Tag

	if err := r.DB.WithContext(ctx).First(&tag, tagID).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r *tagRepository) ReadAll(ctx context.Context) ([]Tag, error) {
	var tags []Tag

	if err := r.DB.WithContext(ctx).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *tagRepository) Delete(ctx context.Context, tagID int) error {
	var tag Tag

	db := r.DB.WithContext(ctx)

	if err := db.First(&tag, tagID).Error; err != nil {
		return err
	}

	if err := db.Delete(&tag).Error; err != nil {
		return err
	}

	return nil
}
