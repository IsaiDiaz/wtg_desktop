package device

type Device struct {
	ID         uint   `gorm:"column:ca_device_id;primaryKey;autoIncrement" json:"id"`
	UUID       string `gorm:"column:ca_device_uuid" json:"uuid"`
	EmployeeID uint   `gorm:"column:ca_employee_id" json:"employee_id"`
}

func (Device) TableName() string {
	return "ca_device"
}
