package category

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(category *Category) error
	GetAll() ([]*Category, error)
	GetByID(id uint) (*Category, error)
	Update(category *Category) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(category *Category) error {
	return r.db.Create(category).Error
}

func (r *repository) GetAll() ([]*Category, error) {
	var categories []*Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *repository) GetByID(id uint) (*Category, error) {
	var category Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *repository) Update(category *Category) error {
	return r.db.Save(category).Error
}

func (r *repository) Delete(id uint) error {
	var category Category
	if err := r.db.First(&category, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&category).Error
}
