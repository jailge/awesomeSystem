package routers

import (
	"awesomeSystem/app/service"
	"awesomeSystem/middleware"
	"awesomeSystem/utils/Commfunction"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var (
	R *gin.Engine
)

func init() {
	R = gin.Default()
	R.Use(cors())
	//R.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	R.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"code": 400, "message": "Bad Request"})
	})
	api()
}
func api() {
	//func New() *Viper
	//New returns an initialized Viper instance.
	//用来生成一个新的 viper
	v := viper.New()
	//对 viper 进行配置
	Commfunction.ConfViper(v)

	ad := middleware.AuthMiddleware(v.GetString("realm"), v.GetString("key"), v.GetString("tokenLookup"), v.GetString("tokenHeadName"))
	authMiddleware, err := jwt.New(ad)
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	R.POST("/login", ad.LoginHandler)
	R.POST("/register", service.AddNewUser)

	R.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth1 := R.Group("/auth")
	// Refresh time can be longer than token timeout
	auth1.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth1.Use(authMiddleware.MiddlewareFunc())
	{
		auth1.GET("/hello", service.HelloHandler)
	}

	acs := R.Group("/api")
	acs.Use(authMiddleware.MiddlewareFunc())
	{
		// 添加一条Policy策略
		acs.POST("acs", service.AddPolicy)
		// 删除一条Policy策略
		acs.DELETE("acs/:id", service.DeletePolicy)
		//// 获取路由列表
		//acs.POST("/routers", middleware.Privilege(), func(c *gin.Context) {
		//	type data struct {
		//		Method string `json:"method"`
		//		Path   string `json:"path"`
		//	}
		//	var datas []data
		//	routers := R.Routes()
		//	for _, v := range routers {
		//		var temp data
		//		temp.Method = v.Method
		//		temp.Path = v.Path
		//		datas = append(datas, temp)
		//	}
		//	APIResponse.C = c
		//	APIResponse.Success(datas)
		//	return
		//})
	}
	// 定义路由组
	user := R.Group("/api")
	// 使用访问控制中间件
	user.Use(middleware.Privilege())
	{
		user.POST("user", func(c *gin.Context) {
			c.JSON(200, gin.H{"code": 200, "message": "user add success"})
		})
		user.DELETE("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user delete success " + id})
		})
		user.PUT("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user update success " + id})
		})
		user.GET("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user Get success " + id})
		})
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
