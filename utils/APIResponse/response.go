package APIResponse

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var C *gin.Context

func Error(message string) {
	if len(message) == 0 {
		message = "fail"
	}
	C.JSON(200, Response{
		Code:    -1,
		Message: message,
		Data:    nil,
	})
}

// Err 返回失败
func Err(c *gin.Context, httpCode int, code int, msg string, jsonStr interface{}) {
	zap.L().Info(msg, zap.Any("调用 Service", fmt.Sprintf("%s 处理出错", msg)), zap.Any("返回错误", jsonStr))
	c.JSON(httpCode, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": jsonStr,
	})
	return
}

func Success(c *gin.Context, code int, msg interface{}, data interface{}) {
	zap.L().Info(msg.(string), zap.Any("调用 Service", fmt.Sprintf("%s 处理请求", msg)), zap.Any("处理返回值", data))
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
