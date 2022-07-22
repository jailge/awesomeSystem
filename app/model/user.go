package model

type User struct {
	//gorm.Model        //加入此行用于在数据库中创建记录的 mate 数据
	UserId   int64  `gorm:"not null"`
	UserName string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100)"`
}

type Permission struct {
	Id       int64  `gorm:"not null"`
	Resource string `gorm:"type:varchar(100)"`
	Action   string `gorm:"type:varchar(100)"`
}

type Role struct {
	Id       int64  `gorm:"not null"`
	RoleName string `gorm:"type:varchar(100)"`
}

type RolePermission struct {
	Id           int64 `gorm:"not null"`
	RoleId       int64 `gorm:"not null"`
	PermissionId int64 `gorm:"not null"`
}

type UserRole struct {
	Id     int64 `gorm:"not null"`
	UserId int64 `gorm:"not null"`
	RoleId int64 `gorm:"not null"`
}
