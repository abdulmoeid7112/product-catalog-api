package models

import (
	"github.com/shopspring/decimal"
)

// Product represents a product in the catalog
// It includes a unique code, price, associated category, and variants
type Product struct {
	ID         uint            `gorm:"primaryKey"`
	Code       string          `gorm:"uniqueIndex;not null"`
	Price      decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	CategoryID uint
	Category   Category
	Variants   []Variant `gorm:"foreignKey:ProductID"`
}

// TableName specifies the database table name for Product
func (p *Product) TableName() string {
	return "products"
}
