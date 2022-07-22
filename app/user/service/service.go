package service

import (
	"awesomeSystem/app/model"
	"awesomeSystem/app/user/Controller"
	"awesomeSystem/app/user/dao"
	"awesomeSystem/app/user/entity"
	"awesomeSystem/global"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strings"

	//"awesomeSystem/app/user/service"
	"awesomeSystem/utils/ACS"
	"awesomeSystem/utils/APIResponse"
	"awesomeSystem/utils/Commfunction"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//var IdentityKey = "id"

func HelloHandler1(c *gin.Context) {
	zap.L().Info("this is hello func", zap.String("test", "test--------->"))
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(global.IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims[global.IdentityKey],
		"userName": user.(*model.User).UserName,
		"text":     "Hello World.",
	})
}

func HelloHandler(c *gin.Context) {
	zap.L().Info("this is hello func", zap.String("test", "test--------->"))
	c.JSON(200, gin.H{
		"text": "Hello World.",
	})
}

func AddNewUser(c *gin.Context) {
	zap.L().Info("AddNewUser", zap.Any("调用 Service", "AddNewUser 处理请求"))
	//user := entity.NewUser{}
	var user entity.NewUser
	err := c.BindJSON(&user)
	if err != nil {
		//zap.L().Info("AddNewUser", zap.Any("返回错误", err.Error()))
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddNewUser error", "")
	}
	//fmt.Println(user)
	hash, _ := Commfunction.PasswordHash(user.Password)

	result := Controller.NewDbUser(entity.NewUser{
		UserName: user.UserName,
		Password: hash,
		Email:    user.Email,
	})
	if result == 0 {
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddNewUser error", result)
	} else {
		APIResponse.Success(c, 200, "新增成功", result)
	}
	//c.JSON(200, gin.H{
	//	"userName": user.UserName,
	//	"result":     result,
	//})
}

func NewUser(c *gin.Context) {
	zap.L().Info("NewUser", zap.Any("调用 Service", "NewUser 处理请求"))
	//user := entity.NewUser{}
	var user entity.NewUser
	err := c.BindJSON(&user)
	if err != nil {
		//zap.L().Info("AddNewUser", zap.Any("返回错误", err.Error()))
		APIResponse.Err(c, http.StatusBadRequest, 400, "NewUser error", "")
	}
	//fmt.Println(user)
	hash, _ := Commfunction.PasswordHash(user.Password)
	//fmt.Println(hash)

	result := Controller.NewDbUser(entity.NewUser{
		UserName: user.UserName,
		Password: hash,
		Email:    user.Email,
	})
	if result == 0 {
		APIResponse.Err(c, http.StatusBadRequest, 400, "NewUser error", result)
	} else {
		APIResponse.Success(c, 200, "新增成功", result)
	}
}

func AddPolicy(c *gin.Context) {
	newPolicy := entity.Policy{}
	err := c.BindJSON(&newPolicy)
	if err != nil {
		return
	}
	//subject := "tom"
	//object := "/api/routers"
	//action := "POST"
	//cacheName := newPolicy.Subject + newPolicy.Object + newPolicy.Action
	result, _ := ACS.Enforcer.AddPolicy(newPolicy.Subject, newPolicy.Object, newPolicy.Action)
	if result {
		// 清除缓存
		//_ = Cache.GlobalCache.Delete(cacheName)
		APIResponse.Success(c, 200, "add Policy success", "")
	} else {
		APIResponse.Err(c, http.StatusBadRequest, 400, "add Policy failed", "")
	}
}

func DeletePolicy(c *gin.Context) {
	policy := entity.Policy{}
	err := c.BindJSON(&policy)
	if err != nil {
		return
	}
	result, _ := ACS.Enforcer.RemovePolicy(policy.Subject, policy.Object, policy.Action)
	if result {
		// 清除缓存 代码省略
		APIResponse.Success(c, 200, "delete Policy success", "")
	} else {
		APIResponse.Err(c, http.StatusBadRequest, 400, "delete Policy failed", "")
	}
}

func GetUserAuth(c *gin.Context) {
	user := c.Param("user")
	per := ACS.Enforcer.GetFilteredNamedPolicy("p", 0, user)
	//fmt.Println(per)

	res := make(map[string][]string)
	for _, v := range per {
		//fmt.Println(k, v)
		//fmt.Println(res)
		pSplit := strings.Split(v[1], "/")
		//fmt.Println("pSplit", pSplit)
		//fmt.Println("pSplit-1", pSplit[1])
		if value, ok := res[pSplit[1]]; ok {
			//fmt.Println("value", value)
			//fmt.Println("ok", ok)
			value = append(value, fmt.Sprintf("%s %s", v[1], v[2]))
			//fmt.Println("value", value)
			res[pSplit[1]] = value
		} else {
			if pSplit[1] == "user" {
				l := []string{fmt.Sprintf("%s %s", v[1], v[2])}
				res[pSplit[1]] = l
			}
		}
		//fmt.Println(res)
		//fmt.Println("****************")
	}
	//fmt.Println(res)
	APIResponse.Success(c, 200, "用户权限列表", res)

}

func GetUserAuthV2(c *gin.Context) {
	user := c.Param("user")
	u, ok := dao.FindUserIdWithName(user)
	if !ok {
		APIResponse.Err(c, http.StatusNotFound, 404, "no user", "")
		return
	}
	ur, ok := dao.FindRoleWithUserId(u.UserId)
	if !ok {
		APIResponse.Success(c, http.StatusOK, "no role", "")
		return
	}
	per, ok := dao.FindPermissionsWithRoleId(ur.RoleId)
	if !ok {
		APIResponse.Success(c, http.StatusOK, "no permission", "")
		return
	}
	res := make(map[string][]string)
	for _, v := range per {
		//fmt.Println(k, v)
		//fmt.Println(res)
		pSplit := strings.Split(v.Resource, "/")
		//fmt.Println("pSplit", pSplit)
		//fmt.Println("pSplit-1", pSplit[1])
		if value, ok := res[pSplit[1]]; ok {
			//fmt.Println("value", value)
			//fmt.Println("ok", ok)
			value = append(value, fmt.Sprintf("%s %s", v.Resource, v.Action))
			//fmt.Println("value", value)
			res[pSplit[1]] = value
		} else {
			l := []string{fmt.Sprintf("%s %s", v.Resource, v.Action)}
			res[pSplit[1]] = l
			//if pSplit[1] == "user" {
			//	l := []string{fmt.Sprintf("%s %s", v.Resource, v.Action)}
			//	res[pSplit[1]] = l
			//}
		}
	}
	APIResponse.Success(c, 200, "用户权限列表", res)
}
