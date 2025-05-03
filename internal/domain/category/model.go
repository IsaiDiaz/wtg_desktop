package category

type Category struct {
	ID      uint    `gorm:"column:ca_category_id;primaryKey;autoIncrement" json:"id"`
	Name    string  `gorm:"column:ca_category_name" json:"name"`
	Payment float64 `gorm:"column:ca_category_payment" json:"payment"`
}

func (Category) TableName() string {
	return "ca_category"
}
