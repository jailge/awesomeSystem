package initialize

import (
	"awesomeSystem/utils/ACS"
	"awesomeSystem/utils/DB/mongodb"
	"awesomeSystem/utils/DB/mysqldb"
	"awesomeSystem/utils/DB/redisdb"
)

// InitMongodb 初始化Mongodb
func InitMongodb() {
	mongodb.Init()
}

// InitRedis 初始化Redis
func InitRedis() {
	redisdb.Init()
}

func InitMySQL() {
	mysqldb.Init()
}

func InitACS() {
	ACS.Init()
}
