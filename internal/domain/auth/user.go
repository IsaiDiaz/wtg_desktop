package auth

type AuthUser struct {
	ID     int    `gorm:"column:auth_user_id;primaryKey;autoIncrement"`
	Name   string `gorm:"column:auth_user_name;size:50;not null"`
	Email  string `gorm:"column:auth_user_email;size:100;not null"`
	Secret string `gorm:"column:auth_user_secret;size:100;not null"`

	UserRoles []AuthUserRole `gorm:"foreignKey:UserID"`
}

func (AuthUser) TableName() string {
	return "auth_user"
}
