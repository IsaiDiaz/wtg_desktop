package auth

type AuthPermission struct {
	ID   int    `gorm:"column:auth_permission_id;primaryKey;autoIncrement"`
	Name string `gorm:"column:auth_permission_name;size:100;not null"`

	RolePermissions []AuthRolePermission `gorm:"foreignKey:PermissionID"`
}

func (AuthPermission) TableName() string {
	return "auth_permission"
}
