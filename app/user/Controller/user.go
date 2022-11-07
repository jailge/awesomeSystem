package Controller

import (
	"awesomeSystem/app/model"
	"awesomeSystem/app/user/dao"
	"awesomeSystem/app/user/entity"
	"awesomeSystem/utils/APIResponse"
	"awesomeSystem/utils/Token"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func HandleUserModelToMap(user *model.User) map[string]interface{} {
	userItemMap := map[string]interface{}{
		"email": user.Email,
		//"nick_name": user.NickName,
		//"head_url":  user.HeadUrl,
		//"birthday":  birthday,
		//"address":   user.Address,
		//"desc":      user.Desc,
		//"gender":    user.Gender,
		//"role":      user.Role,
		//"mobile":    user.Mobile,
	}
	return userItemMap
}

// PasswordLogin 登录
func PasswordLogin(c *gin.Context) {
	var user entity.Login
	err := c.BindJSON(&user)
	if err != nil {
		//zap.L().Info("AddNewUser", zap.Any("返回错误", err.Error()))
		APIResponse.Err(c, http.StatusBadRequest, 400, "user error", "")
	}
	//PasswordLoginForm := forms.PasswordLoginForm{
	//
	//}
	//if err := c.ShouldBind(&PasswordLoginForm); err != nil {
	//	utils.HandleValidatorError(c, err)
	//	return
	//}
	//// 数字验证码验证失败    store.Verify(验证码id,验证码,验证后是否关闭)
	//if !store.Verify(PasswordLoginForm.CaptchaId, PasswordLoginForm.Captcha, true) {
	//	Response.Err(c, 400, 400, "验证码错误", "")
	//	return
	//}
	//查询数据库是否有用户
	//var exist model.User
	exist, ok := dao.FindUserInfo(user.UserName, user.Password)
	if !ok {
		APIResponse.Err(c, 401, 401, "用户账号密码错误", "")
		return
	}
	//
	ur, ok := dao.FindRoleWithUserId(exist.UserId)
	var roleId int64
	if !ok {
		roleId = 0
	} else {
		roleId = ur.RoleId
	}
	token := Token.CreateToken(c, exist.UserId, exist.UserName, roleId)
	userinfoMap := HandleUserModelToMap(exist)
	userinfoMap["token"] = token
	//APIResponse.Success(c, 200, "success", userinfoMap)
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg":  "登录成功",
		"data": token,
	})
}

func AddRole(c *gin.Context) {
	zap.L().Info("AddRole", zap.Any("调用 Service", "AddRole 处理请求"))
	newRole := entity.NewRole{}
	err := c.BindJSON(&newRole)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddRole 参数错误", newRole.RoleName)
		return
	}
	role, ok := dao.NewRole(newRole.RoleName)
	if !ok {
		APIResponse.Err(c, http.StatusBadRequest, 400, "add role 错误", newRole.RoleName)
		return
	}
	APIResponse.Success(c, http.StatusOK, "add role success", role)
}

func AllRoles(c *gin.Context) {
	zap.L().Info("AllRoles", zap.Any("调用 Service", "AllRoles 处理请求"))
	//var roles []model.Role
	roles, ok := dao.AllRoles()
	if !ok {
		APIResponse.Err(c, http.StatusBadRequest, 400, "all role 错误", roles)
		return
	}
	APIResponse.Success(c, http.StatusOK, "all role success", roles)
}

func AddRolePermission(c *gin.Context) {
	zap.L().Info("AddRolePermission", zap.Any("调用 Service", "AddRolePermission 处理请求"))
	newRolePer := entity.NewRolePermission{}
	err := c.BindJSON(&newRolePer)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddRolePermission 参数错误", newRolePer)
		return
	}
	rolePer, ok := dao.NewRolePermission(newRolePer.RoleId, newRolePer.PermissionId)
	if !ok {
		APIResponse.Err(c, http.StatusBadRequest, 400, "add role 错误", newRolePer)
		return
	}
	APIResponse.Success(c, http.StatusOK, "add role success", rolePer)
}

func AllPermissions(c *gin.Context) {
	zap.L().Info("AllPermissions", zap.Any("调用 Service", "AllPermissions 处理请求"))
	//var roles []model.Role
	pers, ok := dao.AllPermissions()
	if !ok {
		APIResponse.Err(c, http.StatusBadRequest, 400, "all permissions 错误", pers)
		return
	}
	APIResponse.Success(c, http.StatusOK, "all role success", pers)
}

func GetPermissionsWithRoleId(c *gin.Context) {
	zap.L().Info("GetPermissionsWithRoleId", zap.Any("调用 Service", "GetPermissionsWithRoleId 处理请求"))
	id := c.Param("id")
	rId, _ := strconv.ParseInt(id, 10, 64)
	rp, ok := dao.FindPermissionsWithRoleId(rId)
	if !ok {
		APIResponse.Err(c, http.StatusBadRequest, 400, "role permissions 错误", rp)
		return
	}
	APIResponse.Success(c, http.StatusOK, fmt.Sprintf("role %d all permissions", rId), rp)
}

func AddUserRole(c *gin.Context) {
	zap.L().Info("AddUserRole", zap.Any("调用 Service", "AddUserRole 处理请求"))
	//id := c.Param("id")
	//uId, _ := strconv.ParseInt(id, 10, 64)
	newUserRole := entity.NewUserRole{}
	err := c.BindJSON(&newUserRole)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddUserRole 参数错误", newUserRole)
		return
	}
	userRole, ok := dao.NewUserRole(newUserRole.UserId, newUserRole.RoleId)
	if !ok {
		APIResponse.Err(c, http.StatusBadRequest, 400, "add user role 错误", userRole)
		return
	}
	APIResponse.Success(c, http.StatusOK, "add user role success", userRole)
}
