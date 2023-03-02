package utils

import (
	"context"
	"time"

	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
)

// 官方文档建议定义全局化变量
var RedisCache cache.Cache

func initRedis() error {
	redisConn, err := config.String("sqlconn")
	if err != nil {
		logs.Info("ERROR: initMySQL error:", err)
		return err
	}
	RedisCache, err = cache.NewCache("redis", `{"conn":"`+redisConn+`"}`)
	if err != nil {
		logs.Info("ERROR: initMySQL error:", err)
		return err
	}

	return nil
}

func SetKey(key string, value interface{}, expireTime time.Duration) error {
	err := RedisCache.Put(context.TODO(), key, value, expireTime)
	if err != nil {
		logs.Info("ERROR: Redis Set Key error:", err, "  key:", key)
	}
	return err
}
