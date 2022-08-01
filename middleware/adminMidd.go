package middleware

import (
	"awesomeSystem/app/user/dao"
	"awesomeSystem/utils/APIResponse"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// IsAdminAuth 判断权限
func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token信息
		claims, _ := c.Get("claims")
		// 获取现在用户信息
		currentUser := claims.(*CustomClaims)
		//fmt.Println(currentUser.AuthorityId)

		// 判断role权限
		rp, ok := dao.FindPermissionsWithRoleId(int64(currentUser.AuthorityId))
		//fmt.Println(rp, ok)
		if !ok {
			APIResponse.Err(c, http.StatusForbidden, 403, "用户没有权限", "")
			//中断下面中间件
			c.Abort()
			return
		}
		path := c.Request.URL.String()
		method := c.Request.Method
		tmpPath := strings.Split(path, "/")
		//fmt.Println(tmpPath)
		//if tmpPath[len(tmpPath)-1] == "*" {
		//	tmpPath = tmpPath[:len(tmpPath)-1]
		//	path = strings.Join(tmpPath, "/")
		//}
		if len(tmpPath) > 3 {
			tmpPath = tmpPath[:len(tmpPath)-1]
			path = strings.Join(tmpPath, "/")
		}
		//fmt.Println(path, method)
		for _, v := range rp {
			if v.Resource == path && v.Action == method {
				c.Next()
				return
			}
		}
		APIResponse.Err(c, http.StatusForbidden, 403, "用户没有权限", "")
		//中断下面中间件
		c.Abort()
		return
		//if currentUser.AuthorityId != 1 {
		//	ctx.JSON(http.StatusForbidden, gin.H{
		//		"msg": "用户没有权限",
		//	})
		//	//中断下面中间件
		//	c.Abort()
		//	return
		//}
		////继续执行下面中间件
		//c.Next()
	}
}
