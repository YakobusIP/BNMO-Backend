package ratescache

import (
	"context"
	"fmt"
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
	setter := client.Set(ctx, key, value, cache.expires*time.Second)
	if setter.Err() != nil {
		fmt.Println("Failed to set value in redis cache")
		return
	}
	fmt.Println("Successfully set rates in redis cache")
}

// Known bugs: connecting to redis cache may fail and display an error
// i/o timeout.
// Possible fix: Restart device
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