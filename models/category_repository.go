package models

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]Category, error)
	Create(ctx context.Context, category *Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

// GetAll returns all categories
func (r *categoryRepository) GetAll(ctx context.Context) ([]Category, error) {
	var categories []Category
	if err := r.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// Create inserts a new category
func (r *categoryRepository) Create(ctx context.Context, category *Category) error {
	if category.Code == "" || category.Name == "" {
		return errors.New("category code and name are required")
	}
	return r.db.WithContext(ctx).Create(category).Error
}
