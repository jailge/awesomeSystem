package api

import (
	"awesomeSystem/app/weight/Service"
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

	R.GET("/permission/:user", Service.GetUserAuth)

	R.POST("/cal/exist_record", Service.ExistCalRecord)

	weight := R.Group("/weight")
	weight.Use(middleware.RateLimitMiddleware(time.Second, 100, 100), middleware.GinLogger(), middleware.GinRecovery(true), middleware.JWTAuth(), middleware.IsAdminAuth())
	{

		weight.GET("/craft", Service.GetAllCraft)
		weight.POST("/craft", Service.AddCraft)
		weight.DELETE("/craft/:id", Service.DeleteCraftWithId)
		weight.PUT("/craft/:id", Service.UpdateCraft)

		weight.GET("/process", Service.GetAllProcess)
		weight.POST("/process", Service.AddProcess)
		weight.DELETE("/process/:id", Service.DeleteProcessWithId)
		//weight.DELETE("/process", Service.DeleteProcessWithId)
		weight.PUT("/process/:id", Service.UpdateProcess)

		weight.GET("/texture", Service.GetAllTexture)
		weight.POST("/texture", Service.AddTexture)
		weight.DELETE("/texture/:id", Service.DeleteTextureWithId)
		weight.PUT("/texture/:id", Service.UpdateTexture)

		weight.GET("/purchase_status", Service.GetAllPurchaseStatus)
		weight.POST("/purchase_status", Service.AddPurchaseStatus)
		weight.DELETE("/purchase_status/:id", Service.DeletePurchaseStatusWithId)
		weight.PUT("/purchase_status/:id", Service.UpdatePurchaseStatus)

		weight.POST("/weigh_multi_condition_search", Service.WeighMultiConditionSearch)
		weight.POST("/cal_multi_condition_search", Service.CalMultiConditionSearch)

		// 添加一条Policy策略
		weight.POST("acs", Service.AddPolicy)
		// 删除一条Policy策略
		weight.DELETE("acs/:id", Service.DeletePolicy)

	}
}
