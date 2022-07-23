package controller

import (
	"net/http"

	"github.com/YakobusIP/BNMO-Backend.git/exchange"
	"github.com/YakobusIP/BNMO-Backend.git/symbolcache"
	"github.com/gin-gonic/gin"
)


var (
	redis symbolcache.SymbolCache = symbolcache.NewSymbolRedisCache("localhost:6379", 0, 86400)
)

func GetSymbolsFromRedis(c *gin.Context) {
	// Pick one of the currencies because 
	

	//var cacheEntry interface{} = redis.GetSymbol("symbols")
	// If cache hit
	/* if cacheEntry != nil {
		fmt.Println("Not nil")
		c.JSON(http.StatusOK, gin.H{"source":"cache",
		"symbols": cacheEntry})
		return
	} */

	// If cache miss
	var symbols interface{} = exchange.GetSymbolsFromAPI().Symbols
	redis.SetSymbol("symbols", symbols)
	c.JSON(http.StatusOK, gin.H{"source":"API", "symbols": symbols})
}