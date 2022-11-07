package api

import (
	"awesomeSystem/app/user/Controller"
	"awesomeSystem/app/user/service"
	"awesomeSystem/global"
	"awesomeSystem/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"time"
)

func Api(R *gin.Engine) {
	//func New() *Viper
	//New returns an initialized Viper instance.
	//用来生成一个新的 viper
	//v := viper.New()
	////对 viper 进行配置
	//Commfunction.ConfViper(v)

	R.Use(ginprom.PromMiddleware(nil))
	// register the `/metrics` route.
	R.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))

	ad := middleware.AuthMiddleware(global.Settings.JwtInfo.Realm, global.Settings.JwtInfo.Key, global.Settings.JwtInfo.TokenLookup, global.Settings.JwtInfo.TokenHeadName)
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
	R.POST("/v2/login", Controller.PasswordLogin)
	//R.POST("/register", service.AddNewUser)
	R.GET("/permission/:user", service.GetUserAuth)
	R.GET("/v2/permission/:user", service.GetUserAuthV2)

	user := R.Group("/user")
	user.GET("/refresh_token", authMiddleware.RefreshHandler)
	user.Use(middleware.RateLimitMiddleware(time.Second, 100, 100), middleware.GinLogger(), middleware.GinRecovery(true), middleware.JWTAuth(), middleware.IsAdminAuth())
	{
		//user.GET("/hello", service.HelloHandler)
		//user.POST("/register", service.AddNewUser)

		user.POST("/add_user", service.NewUser)

		//user.GET("/permission", service.GetUserAuth)
		// 添加一条Policy策略
		//user.POST("acs", service.AddPolicy)
		// 删除一条Policy策略
		//user.DELETE("acs/:id", service.DeletePolicy)
		user.POST("/user_role", Controller.AddUserRole)
	}

	auth := R.Group("/auth")
	//auth.Use(middleware.JWTAuth(), middleware.IsAdminAuth())
	auth.Use(middleware.RateLimitMiddleware(time.Second, 100, 100), middleware.GinLogger(), middleware.GinRecovery(true), middleware.JWTAuth(), middleware.IsAdminAuth())
	{
		auth.POST("/role", Controller.AddRole)
		auth.GET("/roles", Controller.AllRoles)

		auth.GET("/permissions", Controller.AllPermissions)

		auth.GET("/role_permissions/:id", Controller.GetPermissionsWithRoleId)

		auth.POST("/role_permission", Controller.AddRolePermission)
	}

	//test := R.Group("/test")
	//test.Use(middleware.JWTAuth(), middleware.IsAdminAuth())
	//{
	//	test.GET("/hello", service.HelloHandler)
	//}

	R.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	//auth1 := R.Group("/auth")
	//// Refresh time can be longer than token timeout
	//auth1.GET("/refresh_token", authMiddleware.RefreshHandler)
	//auth1.Use(authMiddleware.MiddlewareFunc())
	//{
	//	auth1.GET("/hello", service.HelloHandler)
	//}

	acs := R.Group("/api")
	//acs.Use(authMiddleware.MiddlewareFunc())
	{
		// 添加一条Policy策略
		acs.POST("acs", service.AddPolicy)
		// 删除一条Policy策略
		acs.DELETE("acs/:id", service.DeletePolicy)

	}
	//// 定义路由组
	//user1 := R.Group("/api")
	//// 使用访问控制中间件
	//user1.Use(middleware.Privilege())
	//{
	//	user1.POST("user", func(c *gin.Context) {
	//		c.JSON(200, gin.H{"code": 200, "message": "user add success"})
	//	})
	//	user1.DELETE("user/:id", func(c *gin.Context) {
	//		id := c.Param("id")
	//		c.JSON(200, gin.H{"code": 200, "message": "user delete success " + id})
	//	})
	//	user1.PUT("user/:id", func(c *gin.Context) {
	//		id := c.Param("id")
	//		c.JSON(200, gin.H{"code": 200, "message": "user update success " + id})
	//	})
	//	user1.GET("user/:id", func(c *gin.Context) {
	//		id := c.Param("id")
	//		c.JSON(200, gin.H{"code": 200, "message": "user Get success " + id})
	//	})
	//}
}
