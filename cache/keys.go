package cache

const (
	Prefix                = "blog:" //项目key前缀
	KeySearchKeyWordTimes = "keyword:searchTimes:"
	KeyEssayKeyword       = "essay:keyword:"
	KeyIndex              = "index:"
)

func getRedisKey(key string) string {
	return Prefix + key
}
