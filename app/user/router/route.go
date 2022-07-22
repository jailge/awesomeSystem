package router

import (
	"awesomeSystem/app/user/api"
	"awesomeSystem/utils/Commfunction"
	"github.com/gin-gonic/gin"
)

var (
	R *gin.Engine
)

func init() {
	R = gin.Default()
	R.SetTrustedProxies([]string{"10.10.181.60", "10.10.181.111", "http://10.10.6.51"})

	R.Use(Commfunction.Cors())
	//R.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	R.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"code": 400, "message": "Bad Request"})
	})
	api.Api(R)
}
