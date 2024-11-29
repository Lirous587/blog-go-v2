package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var client *redis.Client

var Rdb *redis.Client

func Init() (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})
	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	Rdb = client
	return nil
}

func Close() {
	if err := client.Close(); err != nil {
		zap.L().Error("redis close() failed", zap.Error(err))
	}
}
