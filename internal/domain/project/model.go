package project

import "time"

type Project struct {
	ID          uint       `gorm:"column:ca_project_id;primaryKey"`
	Name        string     `gorm:"column:ca_project_name"`
	InitialDate time.Time  `gorm:"column:ca_project_initial_date"`
	FinalDate   *time.Time `gorm:"column:ca_project_final_date"`
	IsCurrent   bool       `gorm:"column:ca_project_is_current"`
	Status      bool       `gorm:"column:ca_project_status"`
	Description string     `gorm:"column:ca_project_description"`
}

func (Project) TableName() string {
	return "ca_project"
}
