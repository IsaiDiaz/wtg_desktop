package projectemployee

import "gorm.io/gorm"

type Repository interface {
	Create(projectEmployee *ProjectEmployee) error
	GetByID(id int) (*ProjectEmployee, error)
	GetByProjectID(projectID int) ([]*ProjectEmployee, error)
	GetByEmployeeID(employeeID int) ([]*ProjectEmployee, error)
	GetAll() ([]*ProjectEmployee, error)
	Update(projectEmployee *ProjectEmployee) error
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(projectEmployee *ProjectEmployee) error {
	return r.db.Create(projectEmployee).Error
}

func (r *repository) GetByID(id int) (*ProjectEmployee, error) {
	var projectEmployee ProjectEmployee
	if err := r.db.First(&projectEmployee, id).Error; err != nil {
		return nil, err
	}
	return &projectEmployee, nil
}

func (r *repository) GetAll() ([]*ProjectEmployee, error) {
	var projectEmployees []*ProjectEmployee
	if err := r.db.Find(&projectEmployees).Error; err != nil {
		return nil, err
	}
	return projectEmployees, nil
}

func (r *repository) Update(projectEmployee *ProjectEmployee) error {
	return r.db.Save(projectEmployee).Error
}

func (r *repository) Delete(id int) error {
	var projectEmployee ProjectEmployee
	if err := r.db.First(&projectEmployee, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&projectEmployee).Error
}

func (r *repository) GetByProjectID(projectID int) ([]*ProjectEmployee, error) {
	var projectEmployees []*ProjectEmployee
	if err := r.db.Where("project_id = ?", projectID).Find(&projectEmployees).Error; err != nil {
		return nil, err
	}
	return projectEmployees, nil
}

func (r *repository) GetByEmployeeID(employeeID int) ([]*ProjectEmployee, error) {
	var projectEmployees []*ProjectEmployee
	if err := r.db.Where("employee_id = ?", employeeID).Find(&projectEmployees).Error; err != nil {
		return nil, err
	}
	return projectEmployees, nil
}
