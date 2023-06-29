package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/castmetal/golang-api-boilerplate/src/config"

	"github.com/redis/go-redis/v9"
)

type IRedisClient interface {
	GetData(ctx context.Context, key string) (string, error)
	SetData(ctx context.Context, key string, value string, ttl time.Duration) error
	DelAllData(ctx context.Context, keyPattern string) error
	DelData(ctx context.Context, key string) (int64, error)
}

type RedisConn struct {
	Client *redis.Client
}

func NewRedisClient(config config.EnvStruct) IRedisClient {
	port, err := strconv.Atoi(config.Cache.Addr)
	if err != nil {
		port = 6379
	}

	redisDb, err := strconv.Atoi(config.Cache.Database)
	if err != nil {
		redisDb = 0
	}

	options := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Cache.Addr, port),
		DB:   redisDb,
	}

	if config.Cache.Password != " " {
		options.Password = config.Cache.Password
	}

	rdb := redis.NewClient(options)

	var redisClient IRedisClient = &RedisConn{
		Client: rdb,
	}

	return redisClient
}

func (rc *RedisConn) GetData(ctx context.Context, key string) (string, error) {
	return rc.Client.Get(ctx, key).Result()
}

func (rc *RedisConn) SetData(ctx context.Context, key string, value string, ttl time.Duration) error {
	return rc.Client.Set(ctx, key, value, ttl).Err()
}

func (rc *RedisConn) DelData(ctx context.Context, key string) (int64, error) {
	return rc.Client.Del(ctx, key).Result()
}

func (rc *RedisConn) DelAllData(ctx context.Context, keyPattern string) error {
	iter := rc.Client.Scan(ctx, 0, keyPattern+"*", 0).Iterator()
	for iter.Next(ctx) {
		err := rc.Client.Del(ctx, iter.Val()).Err()
		if err != nil {
			return err
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}
