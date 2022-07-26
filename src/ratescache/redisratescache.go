package ratescache

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
)

var (
	ctx = context.Background()
)

type redisRatesCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRatesRedisCache(host string, db int, expires time.Duration) RatesCache {
	return &redisRatesCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (cache *redisRatesCache) getRatesClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisRatesCache) SetRates(key string, value float64) {
	client := cache.getRatesClient()
	client.Set(ctx, key, value, cache.expires*time.Second)
}

func (cache *redisRatesCache) GetRates(key string) float64 {
	client := cache.getRatesClient()

	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return -1
	}

	val_float, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
	if err != nil {
		return -2
	}

	return float64(val_float)
}