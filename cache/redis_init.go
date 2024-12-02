package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var client *redis.Client

func RedisInit() (*redis.Client, error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return client, nil
}

func Close() {
	if err := client.Close(); err != nil {
		zap.L().Error("redis close() failed", zap.Error(err))
	}
}
