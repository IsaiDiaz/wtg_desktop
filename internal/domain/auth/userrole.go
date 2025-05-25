package auth

type AuthUserRole struct {
	ID     int `gorm:"column:auth_user_role_id;primaryKey;autoIncrement"`
	UserID int `gorm:"column:auth_user_id;not null"`
	RoleID int `gorm:"column:auth_role_id;not null"`

	Role *AuthRole `gorm:"foreignKey:RoleID"`
	User *AuthUser `gorm:"foreignKey:UserID"`
}

func (AuthUserRole) TableName() string {
	return "auth_user_role"
}
