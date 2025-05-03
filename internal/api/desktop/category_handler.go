package desktop

import (
	"wtg_desktop/internal/domain/category"
)

type CategoryHandler struct {
	service category.Service
}

func NewCategoryHandler(service category.Service) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (c *CategoryHandler) CreateCategory(cat *category.Category) error {
	return c.service.CreateCategory(cat)
}

func (c *CategoryHandler) GetCategoryByID(id uint) (*category.Category, error) {
	return c.service.GetCategoryByID(id)
}

func (c *CategoryHandler) GetAllCategories() ([]*category.Category, error) {
	return c.service.GetAllCategories()
}

func (c *CategoryHandler) UpdateCategory(cat *category.Category) error {
	return c.service.UpdateCategory(cat)
}

func (c *CategoryHandler) DeleteCategory(id uint) error {
	return c.service.DeleteCategory(id)
}
