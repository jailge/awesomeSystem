package Controller

import (
	"awesomeSystem/app/user/entity"
	"awesomeSystem/utils/DB/mysqldb"
	"fmt"
	"go.uber.org/zap"
)

func NewDbUser(user entity.NewUser) int {
	fmt.Println("****")
	fmt.Println(user)
	result := mysqldb.Mysql.Table("user").Select("user_name", "password", "email").Create(&user)

	if result.Error != nil {
		//log.Fatal("error new user")
		zap.L().Info("db add new user", zap.Any("返回错误", result.Error))
		return 0
	}
	return int(result.RowsAffected)
}
