package models

import (
	"context"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]Product, error)
	List(ctx context.Context, filter ProductFilter) ([]Product, int64, error)
}

type ProductFilter struct {
	CategoryCode string
	MaxPrice     *decimal.Decimal
	Offset       int
	Limit        int
}

type productRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll(ctx context.Context) ([]Product, error) {
	var products []Product
	err := r.db.WithContext(ctx).
		Preload("Variants").
		Preload("Category").
		Find(&products).Error

	return products, err
}

func (r *productRepository) List(ctx context.Context, filter ProductFilter) ([]Product, int64, error) {

	var (
		products []Product
		total    int64
	)

	query := r.db.WithContext(ctx).
		Model(&Product{}).
		Joins("LEFT JOIN categories ON categories.id = products.category_id")

	// Filtering
	if filter.CategoryCode != "" {
		query = query.Where("categories.code = ?", filter.CategoryCode)
	}

	if filter.MaxPrice != nil {
		query = query.Where("products.price < ?", filter.MaxPrice)
	}

	// Count before pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated data
	err := query.
		Preload("Variants").
		Preload("Category").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Find(&products).Error

	return products, total, err
}
