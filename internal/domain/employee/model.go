package employee

import (
	"time"
)

type Employee struct {
	ID         int       `gorm:"column:ca_employee_id;primaryKey;autoIncrement"`
	Name       string    `gorm:"column:ca_employee_name;size:50;not null"`
	CI         string    `gorm:"column:ca_employee_ci;size:30;not null"`
	BirthDate  time.Time `gorm:"column:ca_employee_birth_date;not null"`
	StartDate  time.Time `gorm:"column:ca_employee_start_date;not null"`
	PhotoURL   string    `gorm:"column:ca_employee_photo_url;size:500;not null"`
	Auth       int       `gorm:"column:ca_employee_auth;not null"`
	CategoryID int       `gorm:"column:ca_category_id;not null"`
	Email      string    `gorm:"column:ca_employee_email;size:50;not null"`
	Phone      string    `gorm:"column:ca_employee_phone;size:50"`
}

func (Employee) TableName() string {
	return "ca_employee"
}
