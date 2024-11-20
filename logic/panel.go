package logic

import (
	"blog/dao/redis"
	"blog/models"
)

func GetUserIpCount(data *models.UserIpForSet) (err error) {
	return redis.GetUserIpCount(data)
}

func GetSearchKeywordRank(data *models.RankKindForZset) (err error) {
	//	得到年月日的keywords的zset
	return redis.GetSearchKeywordRank(data)
}
