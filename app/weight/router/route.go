package router

import (
	"awesomeSystem/app/weight/api"
	"awesomeSystem/middleware"
	"github.com/gin-gonic/gin"
	//ginprometheus "github.com/zsais/go-gin-prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	//"github.com/chenjiandongx/ginprom"
)

var (
	R *gin.Engine
)

func init() {
	R = gin.Default()
	R.SetTrustedProxies([]string{"10.10.181.60", "10.10.181.111", "http://10.10.6.51"})

	// Optional custom metrics list
	//customMetrics := []*ginprometheus.Metric{
	//	&ginprometheus.Metric{
	//		ID:          "1234",                // optional string
	//		Name:        "test_metric",         // required string
	//		Description: "Counter test metric", // required string
	//		Type:        "counter",             // required string
	//	},
	//	&ginprometheus.Metric{
	//		ID:          "1235",                // Identifier
	//		Name:        "test_metric_2",       // Metric Name
	//		Description: "Summary test metric", // Help Description
	//		Type:        "summary",             // type associated with prometheus collector
	//	},
	//	// Type Options:
	//	//	counter, counter_vec, gauge, gauge_vec,
	//	//	histogram, histogram_vec, summary, summary_vec
	//}

	//p := ginprometheus.NewPrometheus("gin", customMetrics)

	//p.Use(R)

	//R.Use(ginprom.PromMiddleware(nil))
	//R.Use(Commfunction.Cors())
	R.Use(middleware.Cors())
	//R.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	R.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"code": 400, "message": "Bad Request"})
	})
	api.Api(R)
}
