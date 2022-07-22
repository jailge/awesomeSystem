package mysqldb

import (
	"fmt"
	"go.uber.org/zap"
)

func NewDbUser(user interface{}) int {
	fmt.Println("****")
	fmt.Println(user)
	result := Mysql.Table("user").Select("user_name", "password", "email").Create(&user)

	if result.Error != nil {
		//log.Fatal("error new user")
		zap.L().Info("db add new user", zap.Any("返回错误", result.Error))
		return 0
	}
	return int(result.RowsAffected)
}
