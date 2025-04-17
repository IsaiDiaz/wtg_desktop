package employee

import (
	"time"
)

type Employee struct {
	ID         uint      `gorm:"column:ca_employee_id;primaryKey" json:"id"`
	Name       string    `gorm:"column:ca_employee_name" json:"name"`
	CI         string    `gorm:"column:ca_employee_ci" json:"ci"`
	BirthDate  time.Time `gorm:"column:ca_employee_birth_date"    json:"birth_date"`
	StartDate  time.Time `gorm:"column:ca_employee_start_date"    json:"start_date"`
	PhotoURL   string    `gorm:"column:ca_employee_photo_url"     json:"photo_url"`
	Auth       int       `gorm:"column:ca_employee_auth" json:"auth"`
	CategoryID int       `gorm:"column:ca_category_id" json:"category_id"`
	Email      string    `gorm:"column:ca_employee_email" json:"email"`
}
