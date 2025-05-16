package projectemployee

type ProjectEmployee struct {
	ID                uint `gorm:"column:ca_project_employee_id;primaryKey"`
	ProjectID         uint `gorm:"column:ca_project_id"`
	EmployeeID        uint `gorm:"column:ca_employee_id"`
	CurrentCategoryID uint `gorm:"column:ca_current_category"`
	Status            bool `gorm:"column:ca_project_employee_status"`
}

func (ProjectEmployee) TableName() string {
	return "ca_project_employee"
}
