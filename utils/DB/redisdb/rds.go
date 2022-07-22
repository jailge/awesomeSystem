package redisdb

import (
	"awesomeSystem/global"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"log"
)

//type RedisDrivers struct {
//	RClient *redis.Client
//
//}

var RdsClient *redis.Client

// Init 初始化
func Init() {

	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", conf.Mongodb.Username, conf.Mongodb.Password, conf.Mongodb.Host, conf.Mongodb.Port)
	//fmt.Println(uri)
	RdsClient = Connect()
	//MgoDbName = conf.Mongodb.Database
}

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.Settings.RedisInfo.Host, global.Settings.RedisInfo.Port),
		Password: global.Settings.RedisInfo.Password,
		DB:       global.Settings.RedisInfo.Db,
	})
	// 通过 client.Ping() 来检查是否成功连接到了 redis 服务器
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		//fmt.Println("redis 连接失败")
		zap.L().Info("Failed to Redis!", zap.String("Redis", "--------->failed"))
		return nil
	}
	//fmt.Println(pong, "redis 连接成功！！！")

	//fmt.Println("Connected to MongoDB!")
	zap.L().Info("Redis connect!", zap.String("Redis", "--------->connect"))

	return client
}
