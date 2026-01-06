package models

// Category represents a product category
type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"uniqueIndex;not null"`
	Name string `gorm:"not null"`
}

// TableName specifies the database table name for Category
func (Category) TableName() string {
	return "categories"
}
