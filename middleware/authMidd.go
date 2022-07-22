package middleware

import (
	"awesomeSystem/app/model"
	"awesomeSystem/global"
	"awesomeSystem/utils/Auth"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	//"google.golang.org/grpc/profiling/service"
	"time"
)

// AuthMiddleware 定义一个中间件，用来反馈 jwt 的认证逻辑
//这里将相应的配置直接以变量的方式传递进来了
func AuthMiddleware(v_realm, v_key, v_tokenLookup, v_tokenHeadName string) *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		//Realm name to display to the user. Required.
		//必要项，显示给用户看的域
		Realm: v_realm,
		//Secret key used for signing. Required.
		//用来进行签名的密钥，就是加盐用的
		Key: []byte(v_key),
		//Duration that a jwt token is valid. Optional, defaults to one hour
		//JWT 的有效时间，默认为一小时
		Timeout: time.Duration(time.Hour * 1000),
		// This field allows clients to refresh their token until MaxRefresh has passed.
		// Note that clients can refresh their token in the last moment of MaxRefresh.
		// This means that the maximum validity timespan for a token is MaxRefresh + Timeout.
		// Optional, defaults to 0 meaning not refreshable.
		//最长的刷新时间，用来给客户端自己刷新 token 用的
		MaxRefresh: time.Hour,
		// Callback function that should perform the authentication of the user based on userID and
		// password. Must return true on success, false on failure. Required.
		// Option return user data, if so, user data will be stored in Claim Array.
		//必要项, 这个函数用来判断 User 信息是否合法，如果合法就反馈 true，否则就是 false, 认证的逻辑就在这里
		Authenticator: Auth.UserAuthCallback,
		// Callback function that should perform the authorization of the authenticated user. Called
		// only after an authentication success. Must return true on success, false on failure.
		// Optional, default to success
		//可选项，用来在 Authenticator 认证成功的基础上进一步的检验用户是否有权限，默认为 success
		Authorizator: Auth.UserAuthPrivCallback,
		// User can define own Unauthorized func.
		//可以用来息定义如果认证不成功的的处理函数
		Unauthorized: Auth.UnAuthFunc,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		//这个变量定义了从请求中解析 token 的位置和格式
		TokenLookup: v_tokenLookup,
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		//TokenHeadName 是一个头部信息中的字符串
		TokenHeadName: v_tokenHeadName,
		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		//这个指定了提供当前时间的函数，也可以自定义
		TimeFunc: time.Now,

		IdentityKey: global.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					global.IdentityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				UserName: claims[global.IdentityKey].(string),
			}
		},
	}
}
