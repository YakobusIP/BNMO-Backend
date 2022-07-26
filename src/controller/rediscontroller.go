package controller

import (
	"BNMO/exchange"
	"BNMO/ratescache"
	"fmt"
	"time"
)


var (
	redis ratescache.RatesCache = ratescache.NewRatesRedisCache("localhost:6379", 0, time.Hour * 24)
)

func GetRatesFromRedis(requestedKey string) (string, float64) {
	// Check cache availability
	cacheEntry := redis.GetRates(requestedKey)
	
	// Cache hit events
	if cacheEntry != -1 && cacheEntry != -2 {
		fmt.Println("Value found within redis")
		return "Cache", cacheEntry
	}

	// Cache miss events
	var rates map[string]float64 = exchange.RequestRatesFromAPI().Rates
	var output float64
	for key, value := range rates {
		if key == requestedKey {
			output = value
		}
		redis.SetRates(key, value)
	}
	
	return "API", output
}