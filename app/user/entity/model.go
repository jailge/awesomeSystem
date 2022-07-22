package entity

type Login struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type NewUser struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

type Policy struct {
	Subject string `json:"subject"`
	Object  string `json:"object"`
	Action  string `json:"action"`
}

type CheckUserAuth struct {
	Token  string `json:"token"`
	Object string `json:"object"`
	Action string `json:"action"`
}

type NewRole struct {
	RoleName string `json:"role_name" binding:"required"`
}

type NewRolePermission struct {
	RoleId       int64 `json:"role_id" binding:"required"`
	PermissionId int64 `json:"permission_id" binding:"required"`
}
