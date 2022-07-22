package Commfunction

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"

	//Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats
	//viper 是一个非常强大的用来进行配置管理的包
	"github.com/spf13/viper"
)

const passSalt = "xulei$xulei"

func ConfViper(v *viper.Viper) {
	//func SetConfigType(in string)
	//SetConfigType sets the type of the configuration returned by the remote source, e.g. "json".
	//这里我使用 toml 的格式来填充配置
	v.SetConfigType("toml")

	//定义一个 byte 的数组，用来存储配置
	//这种方式是直接把配置写到内存中
	//在测试环境下和配置比较少的情况下，可以直接使用这种方式来快速实现
	var tomlConf = []byte(`
	port='8080'
	realm='zone name just for test'
	key='secret key salt'
	tokenLookup='header: Authorization, query: token, cookie: jwt'
	tokenHeadName='Bearer'
	loginPath='/login'
	authPath='/auth'
	refreshPath='/refresh_token'
	testPath='/hello'
	db_host     = "10.10.6.90"
	db_port     = "3306"
	db_user     = "root"
	db_name   = "mdm"
	db_password = "root"
	model_config = 'config/model.conf'
	`)
	//func ReadConfig(in io.Reader) error
	//ReadConfig will read a configuration file, setting existing keys to nil if the key does not exist in the file.
	//用来从上面的byte数组中读取配置内容
	v.ReadConfig(bytes.NewBuffer(tomlConf))

	//配置默认值，如果配置内容中没有指定，就使用以下值来作为配置值，给定默认值是一个让程序更健壮的办法
	v.SetDefault("port", "8000")
	v.SetDefault("realm", "testzone")
	v.SetDefault("key", "secret")
	v.SetDefault("tokenLookup", "header: Authorization, query: token, cookie: jwt")
	v.SetDefault("tokenHeadName", "Bearer")
	v.SetDefault("loginPath", "/login")
	v.SetDefault("authPath", "/auth")
	v.SetDefault("refreshPath", "/refresh_token")
	v.SetDefault("testPath", "/hello")
	v.SetDefault("db_host", "127.0.0.1")
	v.SetDefault("db_port", "5432")
	v.SetDefault("db_user", "postgresql")
	v.SetDefault("db_name", "testdb")
	v.SetDefault("db_password", "123456")
	v.SetDefault("admin_name", "admin")
	v.SetDefault("admin_pass", "admin")
	v.SetDefault("casbin_config", "./auth.conf")
	v.SetDefault("casbin_config", "./auth.csv")
}

// Md5Password 密码加密
func Md5Password(pass string) string {
	w := md5.New()

	io.WriteString(w, pass+passSalt)     //将str写入到w中
	return fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
}

func Cors() gin.HandlerFunc {
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

func GetNowFormatTodayTime() string {
	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()),
		now.Day())

	return dateStr
}
