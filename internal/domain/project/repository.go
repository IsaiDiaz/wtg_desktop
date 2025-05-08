package project

import "gorm.io/gorm"

type Repository interface {
	Create(project *Project) error
	FindAll() ([]*Project, error)
	FindByID(id uint) (*Project, error)
	Update(project *Project) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(project *Project) error {
	return r.db.Create(project).Error
}

func (r *repository) FindAll() ([]*Project, error) {
	var projects []*Project
	err := r.db.Find(&projects).Error
	return projects, err
}

func (r *repository) FindByID(id uint) (*Project, error) {
	var project Project
	err := r.db.First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *repository) Update(project *Project) error {
	return r.db.Save(project).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Project{}, id).Error
}
