package device

import "gorm.io/gorm"

type Repository interface {
	Create(device *Device) error
	GetByID(id int) (*Device, error)
	GetAll() ([]*Device, error)
	Update(device *Device) error
	Delete(id int) error
	GetByEmployeeID(employeeID int) ([]*Device, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(device *Device) error {
	return r.db.Create(device).Error
}

func (r *repository) GetByID(id int) (*Device, error) {
	var device Device
	if err := r.db.First(&device, id).Error; err != nil {
		return nil, err
	}
	return &device, nil
}

func (r *repository) GetAll() ([]*Device, error) {
	var devices []*Device
	if err := r.db.Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}

func (r *repository) Update(device *Device) error {
	return r.db.Save(device).Error
}

func (r *repository) Delete(id int) error {
	var device Device
	if err := r.db.First(&device, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&device).Error
}

func (r *repository) GetByEmployeeID(employeeID int) ([]*Device, error) {
	var devices []*Device
	if err := r.db.Where("employee_id = ?", employeeID).Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}
