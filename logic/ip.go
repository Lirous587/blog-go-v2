package logic

import (
	"blog/dao/redis"
)

const (
	preSecondRequest = 15
	maliciousTimes   = 100
)

const (
	requestTooFrequent = "请求过于频繁,请稍后再试"
	ipForbid           = "恭喜你!ip已被永久封禁"
)

func SaveUserIp(ip string) error {
	if err := redis.SaveUserIp(ip); err != nil {
		return err
	}
	return nil
}
