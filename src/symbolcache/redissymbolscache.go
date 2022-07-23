package symbolcache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

var (
	ctx = context.Background()
)

type redisSymbolCache struct {
	host    string
	db      int
	expires time.Duration
}


func NewSymbolRedisCache(host string, db int, expires time.Duration) SymbolCache {
	return &redisSymbolCache{
		host: host,
		db: db,
		expires: expires,
	}
}


func (cache *redisSymbolCache) getSymbolClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cache.host,
		Password: "",
		DB: cache.db,
	})
}

func (cache *redisSymbolCache) SetSymbol(key string, value interface{}) {
	client := cache.getSymbolClient()
	err := client.Set(ctx, key, value, cache.expires*time.Second).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (cache *redisSymbolCache) GetSymbol(key string) interface{} {
	client := cache.getSymbolClient()
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("Key does not exist")
		return nil
	}

	return interface{}(val)
}

