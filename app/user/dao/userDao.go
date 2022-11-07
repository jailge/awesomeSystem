package dao

import (
	"awesomeSystem/app/model"
	"awesomeSystem/utils/Commfunction"
	"awesomeSystem/utils/DB/mysqldb"
)

var user model.User

type rolePermission struct {
	RoleName string
	Resource string
	Action   string
}

// NewRole 新建Role
func NewRole(roleName string) (*model.Role, bool) {
	role := model.Role{RoleName: roleName}
	res := mysqldb.Mysql.Table("role").Select("role_name").Create(&role)
	if res.RowsAffected < 1 {
		return &role, false
	}
	return &role, true
}

func AllRoles() ([]*model.Role, bool) {
	var roles []*model.Role
	res := mysqldb.Mysql.Table("role").Find(&roles)
	if res.RowsAffected < 1 {
		return roles, false
	}
	return roles, true
}

func NewRolePermission(roleId, perId int64) (*model.RolePermission, bool) {
	rp := model.RolePermission{RoleId: roleId, PermissionId: perId}
	res := mysqldb.Mysql.Table("role_permission").Select("role_id", "permission_id").Create(&rp)
	if res.RowsAffected < 1 {
		return &rp, false
	}
	return &rp, true
}

func AllPermissions() ([]*model.Permission, bool) {
	var per []*model.Permission
	res := mysqldb.Mysql.Table("permission").Find(&per)
	if res.RowsAffected < 1 {
		return per, false
	}
	return per, true
}

// FindUserInfo UsernameFindUserInfo 通过username找到用户信息
func FindUserInfo(username string, password string) (*model.User, bool) {
	// 查询用户

	rows := mysqldb.Mysql.Table("user").Where("user_name = ?", username).Find(&user)
	if user.UserName != "" {
		//定义一个临时的结构对象
		queryRes := model.User{} //创建一个临时的存放空间
		//将 user_id 为认证信息中的 密码找出来(目前密码是明文的，这个其实不安全，可以通过加盐哈希将结果进行对比的方式以提高安全等级，这里只作原理演示，就不搞那么复杂了)
		//找到后放到前面定义的临时结构变量里
		mysqldb.Mysql.Table("user").Where("user_name = ?", user.UserName).Find(&queryRes)
		//fmt.Println("*********")
		//fmt.Println(user)
		//fmt.Println(queryRes)

		//对比，如果密码也相同，就代表认证成功了
		//fmt.Println(user.Password, queryRes.Password)
		match := Commfunction.PasswordVerify(password, queryRes.Password)
		if match == true {
			if rows.RowsAffected < 1 {
				return &user, false
			}
			//// 查询
			//mysqldb.Mysql.Table("user_role").Where("user_id", queryRes.UserId)
			return &user, true
		} else {
			return &user, false
		}
	}

	//rows := mysqldb.Mysql.Table("user").Where(&model.User{UserName: username, Password: password}).Find(&user)
	//fmt.Println(&user)
	//if rows.RowsAffected < 1 {
	//	return &user, false
	//}
	return &user, false
}

func FindRoleWithUserId(userId int64) (*model.UserRole, bool) {
	// 查询
	var userRole model.UserRole
	rows := mysqldb.Mysql.Table("user_role").Where("user_id", userId).Find(&userRole)
	if rows.RowsAffected < 1 {
		return &userRole, false
	}
	return &userRole, true
}

func FindPermissionIdWithRoleId(roleId int64) (*model.RolePermission, bool) {
	var rolePer model.RolePermission
	rows := mysqldb.Mysql.Table("role_permission").Where("role_id", roleId).Find(&rolePer)
	if rows.RowsAffected < 1 {
		return &rolePer, false
	}
	return &rolePer, true
}

func FindPermissionsWithRoleId(roleId int64) ([]rolePermission, bool) {
	//fmt.Println(roleId)
	//var r rolePermission
	var rp []rolePermission
	sql := "select role.role_name, permission.resource, permission.action from role, permission, role_permission where role.id = role_permission.role_id and permission.id=role_permission.permission_id and role.id=?"
	rows := mysqldb.Mysql.Raw(sql, roleId).Scan(&rp)
	//fmt.Println(rp)
	if rows.RowsAffected < 1 {
		return rp, false
	}
	return rp, true
}

func FindUserIdWithName(userName string) (*model.User, bool) {
	var user model.User
	rows := mysqldb.Mysql.Table("user").Where("user_name", userName).Find(&user)
	if rows.RowsAffected < 1 {
		return &user, false
	}
	return &user, true
}

func NewUserRole(userId, roleId int64) (*model.UserRole, bool) {
	ur := model.UserRole{UserId: userId, RoleId: roleId}
	res := mysqldb.Mysql.Table("user_role").Select("user_id", "role_id").Create(&ur)
	if res.RowsAffected < 1 {
		return &ur, false
	}
	return &ur, true
}
