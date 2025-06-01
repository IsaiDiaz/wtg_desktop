package employee

import "gorm.io/gorm"

type Repository interface {
	Create(employee *Employee) error
	GetByID(id int) (*Employee, error)
	GetAll() ([]*Employee, error)
	Update(employee *Employee) error
	Delete(id int) error
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

func (r *repository) GetByID(id int) (*Employee, error) {
	var employee Employee
	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *repository) GetAll() ([]*Employee, error) {
	var employees []*Employee
	if err := r.db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *repository) Update(employee *Employee) error {
	return r.db.Save(employee).Error
}

func (r *repository) Delete(id int) error {
	return r.db.Delete(&Employee{}, id).Error
}
