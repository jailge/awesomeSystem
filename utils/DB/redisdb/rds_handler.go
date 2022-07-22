package redisdb

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func GetLock(k string, v string) (bool, error) {
	res, err := RdsClient.SetNX(k, v, time.Minute).Result()
	if err != nil {
		zap.L().Info("GetLock error", zap.String("GetLock error", fmt.Sprintf("--------->%s", zap.Error(err).String)))

	}
	return res, err
}

func RunEvalDel(k string, v string) (int, error) {
	//执行Lua脚本
	//:param key: key
	//:param value: value
	//:return: 删除成功1，失败0
	res, err := RdsClient.Eval("if redis.call('get', KEYS[1])==ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end", []string{k}, v).Result()
	if err != nil {
		zap.L().Info("RunEvalDel error", zap.String("RunEvalDel error", fmt.Sprintf("--------->%s", zap.Error(err).String)))
	}
	r, _ := res.(int)
	return r, err
}

func Pttl(k string) (int64, error) {
	//获取key对应剩余过期时间
	//:param key:
	//:return: ms
	pttl, err := RdsClient.PTTL(k).Result()
	if err != nil {
		zap.L().Info("Pttl error", zap.String("Pttl error", fmt.Sprintf("--------->%s", zap.Error(err).String)))
	}
	return int64(pttl), err
}
