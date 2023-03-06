package utils

import (
	"context"
	"time"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/redis/go-redis/v9"
)

// 用go-redis库可以存储多个数据类型
// 如果用beego自带的cache，redis引擎，只能存string类型
var RedisCache *redis.Client

func initRedis() error {
	redisConn, err := config.String("redishost")

	if err != nil {
		logs.Info("ERROR: redis init error:", err)
	}

	RedisCache = redis.NewClient(&redis.Options{
		Addr:     redisConn,
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	err = RedisCache.Set(context.Background(), "key_string", "test", time.Minute*10).Err()
	// fmt.Println(err)
	return nil
}
