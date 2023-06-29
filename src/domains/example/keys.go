package example

const REDIS_LIST_ALL_EXAMPLES_KEY = "/examples/"

func GetRedisKeys() map[string]string {

	redisKeys := make(map[string]string, 1)
	redisKeys["REDIS_LIST_ALL_EXAMPLES_KEY"] = "/examples/"

	return redisKeys
}
