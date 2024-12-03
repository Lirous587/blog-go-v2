package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type UserCache interface {
	GenToken(uid int64) (string, error)
	ParseToken(token string) (int64, error)
}

type UserCacheRedis struct {
	client *redis.Client
	db     *sqlx.DB
}

func NewUserCacheRedis(client *redis.Client) *UserCacheRedis {
	return &UserCacheRedis{
		client: client,
	}
}

func (r *UserCacheRedis) GenToken(uid int64) (string, error) {
	keyPre := getRedisKey(KeyUserToken)
	key := fmt.Sprintf("%s%d:", keyPre, uid)
	pipe := r.client.Pipeline()
	// 获取当前令牌
	currentToken, err := r.getCurrentToken(key, pipe)
	if err == nil && currentToken != "" {
		// 删除当前令牌
		if err := r.invalidateToken(key, pipe); err != nil {
			return "", err
		}
	}
	// 生成新令牌
	token, err := GenerateSecureToken()
	if err != nil {
		return "", fmt.Errorf("cache GenToken failed, err: %w", err)
	}
	expiration := time.Duration(viper.GetInt("auth.expire_hour")) * time.Hour

	// 保存新令牌
	if err = pipe.Set(key, token, expiration).Err(); err != nil {
		return "", err
	}
	_, err = pipe.Exec()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *UserCacheRedis) getCurrentToken(key string, pipe redis.Pipeliner) (string, error) {
	token, err := pipe.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("cache GetCurrentToken failed, err: %w", err)
	}
	return token, nil
}

func (r *UserCacheRedis) invalidateToken(key string, pipe redis.Pipeliner) error {
	if err := pipe.Del(key).Err(); err != nil {
		return fmt.Errorf("cache InvalidateToken failed, err: %w", err)
	}
	return nil
}

func (r *UserCacheRedis) ParseToken(token string) (int64, error) {
	key := getRedisKey(KeyUserToken)
	uidS, err := r.client.HGet(key, token).Result()
	if err != nil {
		return 0, fmt.Errorf("cache ParseToken failed,err:%w", err)
	}
	uid, err := strconv.Atoi(uidS)
	uid64 := int64(uid)
	return uid64, nil
}
