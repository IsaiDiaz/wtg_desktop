package rfidcardhistory

import "gorm.io/gorm"

type Repository interface {
	Create(record *RfidCardHistory) error
	GetByEmployeeID(employeeID uint) ([]RfidCardHistory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(record *RfidCardHistory) error {
	return r.db.Create(record).Error
}

func (r *repository) GetByEmployeeID(employeeID uint) ([]RfidCardHistory, error) {
	var history []RfidCardHistory
	err := r.db.Where("ca_employee_id = ?", employeeID).Find(&history).Error
	return history, err
}
