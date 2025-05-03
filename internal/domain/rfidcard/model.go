package rfidcard

import "time"

type RfidCard struct {
	ID            uint      `gorm:"column:rfid_card_id;primaryKey;autoIncrement" json:"id"`
	UUID          string    `gorm:"column:rfid_card_uuid" json:"uuid"`
	Status        bool      `gorm:"column:rfid_card_status" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	DeactivatedAt time.Time `gorm:"column:deactivate_at" json:"deactivated_at"`
}

func (RfidCard) TableName() string {
	return "rfid_card"
}
