package mysqldb

import (
	"awesomeSystem/global"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Mysql *gorm.DB
)

func Init() {
	var err error
	//dsn := "root:root@(10.10.6.90:3306)/mdm?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", global.Settings.MysqlInfo.Name, global.Settings.MysqlInfo.Password, global.Settings.MysqlInfo.Host, global.Settings.MysqlInfo.Port, global.Settings.MysqlInfo.DBName)
	//fmt.Println(dsn)

	Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//Mysql, err = gorm.Open("mysql", dsn)
	if err != nil {
		//fmt.Println("connect DB error", err.Error())
		//panic(err)
		zap.L().Info("db init", zap.Any("返回错误", err.Error()))
	}
	zap.L().Info("Mysql connect!", zap.String("Mysql", "--------->connect"))
}
