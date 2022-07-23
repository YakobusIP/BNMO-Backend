package convertcache

import (
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

type redisConvertCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewConvertRedisCache(host string, db int, expires time.Duration) ConvertCache {
	return &redisConvertCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (cache *redisConvertCache) getConvertClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisConvertCache) SetConvert(key string, value int) {
	client := cache.getConvertClient()
	client.Set(ctx, key, value, cache.expires*time.Second)
}

func (cache *redisConvertCache) GetConvert(key string) int {
	client := cache.getConvertClient()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return -1
	}

	val_int, err := strconv.Atoi(val)
	if err != nil {
		return -1
	}

	return val_int
}