package mocks

import (
	"context"

	. "github.com/mytheresa/go-hiring-challenge/models"
)

type MockCategoryRepository struct {
	Categories []Category
	Err        error
}

func (m *MockCategoryRepository) GetAll(ctx context.Context) ([]Category, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Categories, nil
}

func (m *MockCategoryRepository) Create(ctx context.Context, category *Category) error {
	if m.Err != nil {
		return m.Err
	}
	category.ID = uint(len(m.Categories) + 1)
	m.Categories = append(m.Categories, *category)
	return nil
}
