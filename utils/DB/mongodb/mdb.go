package mongodb

import (
	"awesomeSystem/global"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"log"
)

//type MConfig struct {
//	Mongodb struct {
//		Host     string `yaml:"host"`
//		Port     string `yaml:"port"`
//		Username string `yaml:"username"`
//		Password string `yaml:"password"`
//		Database string `yaml:"database"`
//	}
//}

//type MongoDrivers struct {
//	Client *mongo.Client
//}

var MgoClient *mongo.Client
var MgoDbName string

// Init 初始化
func Init() {
	//conf := MConfig{}
	//confFile, err := ioutil.ReadFile(dbConf)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = yaml.Unmarshal(confFile, &conf)
	//if err != nil {
	//	log.Fatal(err)
	//}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", global.Settings.MongodbInfo.Name, global.Settings.MongodbInfo.Password, global.Settings.MongodbInfo.Host, global.Settings.MongodbInfo.Port)
	//fmt.Println(uri)
	MgoClient = Connect(uri)
	MgoDbName = global.Settings.MongodbInfo.DBName
}

func Connect(uri string) *mongo.Client {

	//confFile, err := ioutil.ReadFile(dbConf)
	//checkErr(err)
	//
	//err = yaml.Unmarshal(confFile, &mSvc.conf)
	//checkErr(err)

	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", mSvc.conf.Mongodb.Username, mSvc.conf.Mongodb.Password, mSvc.conf.Mongodb.Host, mSvc.conf.Mongodb.Port)

	// 设置客户端参数
	clientOptions := options.Client().ApplyURI(uri)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	//defer client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// 检查链接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Connected to MongoDB!")

	zap.L().Info("Connected to MongoDB!", zap.String("MongoDB", "--------->connect"))
	return client
}

// Close 关闭
func Close() {

	err := MgoClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	zap.L().Info("MongoDB closed!", zap.String("MongoDB", "--------->close"))
	//fmt.Println("Connection to MongoDB closed.")
}

//type MService struct {
//	client *mongo.Client
//	conf   Config
//}

//func checkErr(err error)  {
//	if err != nil {
//		log.Println(err)
//	}
//}

//func (mSvc *MService)InitMongoDB(dbConf string) error {
//
//	confFile, err := ioutil.ReadFile(dbConf)
//	checkErr(err)
//
//	err = yaml.Unmarshal(confFile, &mSvc.conf)
//	checkErr(err)
//
//	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", mSvc.conf.Mongodb.Username, mSvc.conf.Mongodb.Password, mSvc.conf.Mongodb.Host, mSvc.conf.Mongodb.Port)
//
//
//	// 设置客户端连接配置
//	clientOptions := options.Client().ApplyURI(uri)
//
//	// 连接到MongoDB
//	mSvc.client, err = mongo.Connect(context.TODO(), clientOptions)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 检查连接
//	err = mSvc.client.Ping(context.TODO(), nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Connected to MongoDB!")
//	return err
//}
//
//func (mSvc *MService)CloseMongoDB() error {
//	err := mSvc.client.Disconnect(context.TODO())
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("Connection to MongoDB closed.")
//	return err
//}
