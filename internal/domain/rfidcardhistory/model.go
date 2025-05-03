package rfidcardhistory

import "time"

type RfidCardHistory struct {
	ID           uint      `gorm:"column:rfid_card_history_id;primaryKey;autoIncrement" json:"id"`
	AssignedAt   time.Time `gorm:"column:assigned_at" json:"assigned_at"`
	UnassignedAt time.Time `gorm:"column:unassigned_at" json:"unassigned_at"`
	EmployeeID   uint      `gorm:"column:ca_employee_id" json:"employee_id"`
	RfidCardID   uint      `gorm:"column:rfid_card_id" json:"rfid_card_id"`
}

func (RfidCardHistory) TableName() string {
	return "ca_rfid_card_history"
}
