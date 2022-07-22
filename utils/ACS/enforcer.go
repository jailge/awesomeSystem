package ACS

import (
	"awesomeSystem/utils/DB/mysqldb"
	"go.uber.org/zap"

	//"github.com/casbin/casbin"

	"github.com/casbin/casbin/v2"
	//"github.com/casbin/gorm-adapter"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	//_ "github.com/go-sql-driver/mysql"
)

var Enforcer *casbin.Enforcer

func Init() {
	// mysql 适配器
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", global.Settings.MysqlInfo.Name, global.Settings.MysqlInfo.Password, global.Settings.MysqlInfo.Host, global.Settings.MysqlInfo.Port, global.Settings.MysqlInfo.DBName)
	//fmt.Println(dsn)
	//adapter := gormadapter.NewAdapterByDB(mysqldb.Mysql)
	adapter, _ := gormadapter.NewAdapterByDB(mysqldb.Mysql)
	// 通过mysql适配器新建一个enforcer
	Enforcer, _ = casbin.NewEnforcer("config/model.conf", adapter)

	Enforcer.EnableAutoSave(true)

	// 日志记录
	Enforcer.EnableLog(true)
	zap.L().Info("casbin connect!", zap.String("casbin", "--------->connect"))
}
