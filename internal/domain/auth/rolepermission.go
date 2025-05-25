package auth

type AuthRolePermission struct {
	ID           int `gorm:"column:auth_role_permission_id;primaryKey;autoIncrement"`
	RoleID       int `gorm:"column:auth_role_id;not null"`
	PermissionID int `gorm:"column:auth_permission_id;not null"`

	Role       *AuthRole       `gorm:"foreignKey:RoleID"`
	Permission *AuthPermission `gorm:"foreignKey:PermissionID"`
}

func (AuthRolePermission) TableName() string {
	return "auth_role_permission"
}
