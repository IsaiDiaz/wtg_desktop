package employee

import "gorm.io/gorm"

type Repository interface {
	Create(employee *Employee) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(employee *Employee) error {
	return r.db.Create(employee).Error
}
