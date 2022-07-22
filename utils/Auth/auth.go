package Auth

import (
	"awesomeSystem/app/model"
	"awesomeSystem/app/user/entity"
	"awesomeSystem/global"
	"awesomeSystem/utils/ACS"
	"awesomeSystem/utils/Commfunction"
	"awesomeSystem/utils/DB/mysqldb"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UserAuthCallback 定义一个回调函数，用来决断用户id和密码是否有效
func UserAuthCallback(c *gin.Context) (interface{}, error) {

	//这里的通过从数据库中查询来判断是否为现存用户，生产环境下一般都会使用数据库来存储账号信息，进行检验和判断

	//fmt.Println(c)

	user := entity.Login{} //创建一个临时的存放空间
	if err := c.BindJSON(&user); err != nil {
		//fmt.Println(err)
		return "", jwt.ErrMissingLoginValues
	}
	reqPwd := user.Password
	zap.L().Info("UserAuthCallback", zap.String("UserAuthCallback user_name", fmt.Sprintf("--------->%s", user.UserName)))
	//fmt.Println(user)
	//如果这条记录存在的的情况下
	mysqldb.Mysql.Table("user").Where("user_name = ?", user.UserName).Find(&user)

	//if !mysqldb.Mysql.Table("user").Where("user_name = ?", user.UserName).Find(&user). {
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
		match := Commfunction.PasswordVerify(reqPwd, queryRes.Password)
		//fmt.Println(match)
		if match == true {
			//反馈相关信息和 true 的值，代表成功
			//fmt.Println("password correct")
			zap.L().Info("UserAuthCallback success", zap.String("UserAuthCallback user_name success", fmt.Sprintf("--------->%s", user.UserName)))
			return &model.User{
				UserName: user.UserName,
			}, nil
		}
	}
	//否则返回失败
	zap.L().Info("UserAuthCallback fail", zap.String("UserAuthCallback user_name fail", fmt.Sprintf("--------->%s", user.UserName)))
	return nil, jwt.ErrFailedAuthentication
}

// UserAuthPrivCallback 定义一个回调函数，用来决断用户在认证成功的前提下，是否有权限对资源进行访问
func UserAuthPrivCallback(user interface{}, c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)
	//fmt.Println("claims ****************")
	fmt.Println(claims)
	userData, _ := c.Get(global.IdentityKey)
	user = userData.(*model.User).UserName
	//fmt.Println(user)
	if v, ok := user.(string); ok {
		//如果可以正常取出 user 的值，就使用 casbin 来验证一下是否具备资源的访问权限
		path := c.Request.URL.String()
		method := c.Request.Method
		//cacheName := v + path + method
		//fmt.Println(path, method)
		ul := ACS.Enforcer.GetPermissionsForUser(v)
		fmt.Println(ul)
		b, _ := ACS.Enforcer.Enforce(v, path, method)
		fmt.Println(b)

		return b

		//// 从数据库中读取&判断
		//// 记录日志
		//ACS.Enforcer.EnableLog(true)
		//// 加载策略规则
		//err := ACS.Enforcer.LoadPolicy()
		//if err != nil {
		//	//log.Println("loadPolicy error")
		//	zap.L().Info("loadPolicy error", zap.String("loadPolicy", fmt.Sprintf("--------->%s", err.Error())))
		//	//panic(err)
		//}
		//// 验证策略规则
		//result, err := ACS.Enforcer.Enforce(v, path, method)
		//if err != nil {
		//	zap.L().Info("No permission found")
		//	//APIResponse.Err(c, "No permission found")
		//	//c.Abort()
		//}
		////c.Next()
		////fmt.Println(result)
		//return result

		//// 从缓存中读取&判断
		//entry, err := Cache.GlobalCache.Get(cacheName)
		//if err == nil && entry != nil {
		//	if string(entry) == "true" {
		//		c.Next()
		//	} else {
		//		APIResponse.Error("access denied")
		//		c.Abort()
		//	}
		//} else {
		//	// 从数据库中读取&判断
		//	// 记录日志
		//	ACS.Enforcer.EnableLog(true)
		//	// 加载策略规则
		//	err := ACS.Enforcer.LoadPolicy()
		//	if err != nil {
		//		log.Println("loadPolicy error")
		//		//panic(err)
		//	}
		//	// 验证策略规则
		//	result, err := ACS.Enforcer.EnforceSafe(v, path, method)
		//	if err != nil {
		//		APIResponse.Error("No permission found")
		//		c.Abort()
		//	}
		//	if !result {
		//		// 添加到缓存中
		//		Cache.GlobalCache.Set(cacheName, []byte("false"))
		//		APIResponse.Error("access denied")
		//		c.Abort()
		//	} else {
		//		Cache.GlobalCache.Set(cacheName, []byte("true"))
		//	}
		//	c.Next()
		//	fmt.Println(result)
		//	return result
		//}

	}
	//默认策略是不允许
	return false
}

// UnAuthFunc 定义一个函数用来处理，认证不成功的情况
func UnAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

//func LoginFunc(c *gin.Context, code int, message string, time time.Time)  {
//	user := service.Login{}
//	err := c.BindJSON(&user)
//	if err != nil {
//		return
//	}
//	fmt.Println(user)
//}
