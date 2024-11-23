package logic

import (
	"blog/dao/redis"
)

func SaveUserIp(ip string) error {
	if err := redis.SaveUserIp(ip); err != nil {
		return err
	}
	return nil
}
