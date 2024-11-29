package cache

import (
	"blog/models"
	"fmt"
	"strings"
	"time"
)

const (
	rankCount int64 = 5
)

type RemainingTime struct {
	Year  int64 // 秒
	Month int64 // 秒
	Week  int64 // 秒
}

func getRemainingTime() RemainingTime {
	shanghaiLocation, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		shanghaiLocation = time.Local
	}

	now := time.Now().In(shanghaiLocation)
	currentYear, currentMonth, _ := now.Date()

	// 计算年末时间
	endOfYear := time.Date(currentYear+1, time.January, 1, 0, 0, 0, 0, shanghaiLocation)

	// 计算下个月的第一天
	firstOfNextMonth := time.Date(currentYear, currentMonth+1, 1, 0, 0, 0, 0, shanghaiLocation)

	// 计算本周末时间
	daysUntilEndOfWeek := time.Saturday - now.Weekday()
	if daysUntilEndOfWeek <= 0 {
		daysUntilEndOfWeek += 7
	}
	endOfWeek := now.Add(time.Duration(daysUntilEndOfWeek*24) * time.Hour).Truncate(24 * time.Hour)

	return RemainingTime{
		Year:  int64(endOfYear.Sub(now).Seconds()),
		Month: int64(firstOfNextMonth.Sub(now).Seconds()),
		Week:  int64(endOfWeek.Sub(now).Seconds()),
	}
}

// SetYearMonthWeekTimesZoneForZset 设置年月日相关Zset
func SetYearMonthWeekTimesZoneForZset(preKey string, param string, scoreIncrement float64) (err error) {
	// 1.首先对成员做去空格和小写化
	member := strings.ToLower(strings.TrimSpace(param))

	// 2.给每个成员一个总的统计次数 分别实现 年 月 周 关键词统计
	yearKey := fmt.Sprintf("%s%s:", preKey, year)
	monthKey := fmt.Sprintf("%s%s:", preKey, month)
	weekKey := fmt.Sprintf("%s%s:", preKey, week)

	// 3.得到剩余时间
	remainingTime := getRemainingTime()

	// 4.用集合实现 --> 内置排序
	pipe := client.Pipeline()

	// 年统计
	pipe.ZIncrBy(yearKey, scoreIncrement, member)
	pipe.Expire(yearKey, time.Duration(remainingTime.Year)*time.Second)

	// 月统计
	pipe.ZIncrBy(monthKey, scoreIncrement, member)
	pipe.Expire(monthKey, time.Duration(remainingTime.Month)*time.Second)

	// 周统计
	pipe.ZIncrBy(weekKey, scoreIncrement, member)
	pipe.Expire(weekKey, time.Duration(remainingTime.Week)*time.Second)

	// 执行管道命令
	if _, err = pipe.Exec(); err != nil {
		return fmt.Errorf("failed to increase Zset score: %w", err)
	}
	return nil
}

func GetYearMonthWeekTimesZoneForZsetRank(rankKind *models.RankKindForZset, preKey string) (err error) {
	//	得到年月日的keywords的zset
	yearKey := fmt.Sprintf("%s%s:", preKey, year)
	monthKey := fmt.Sprintf("%s%s:", preKey, month)
	weekKey := fmt.Sprintf("%s%s:", preKey, week)

	// 从每个zset中获取前10条数据
	yearList, err := getTopXFromZSet(yearKey, rankCount)
	if err != nil {
		return err
	}

	monthList, err := getTopXFromZSet(monthKey, rankCount)
	if err != nil {
		return err
	}

	weekList, err := getTopXFromZSet(weekKey, rankCount)
	if err != nil {
		return err
	}

	// 合并结果
	*rankKind = models.RankKindForZset{
		Year:  yearList,
		Month: monthList,
		Week:  weekList,
	}

	return nil
}

// getTopXFromZSet 得到Zset排序
func getTopXFromZSet(key string, count int64) (models.RankListForZset, error) {
	result, err := client.ZRevRangeWithScores(key, 0, count).Result()
	if err != nil {
		return models.RankListForZset{}, err
	}

	var rankList models.RankListForZset
	for _, z := range result {
		rankList.X = append(rankList.X, z.Member.(string))
		rankList.Y = append(rankList.Y, int(z.Score))
	}
	return rankList, nil
}
