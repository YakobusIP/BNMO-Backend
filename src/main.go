package main

import (
	"BNMO/database"
	"BNMO/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
)

func main() {
	// Initialize database using GORM
	database.Initialize()

	// Set up CORS policy
	Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "Static"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	  }))
	routes.MapUrls(Router)
	Router.Run()
}

/* func main() {
	//var symbols exchange.SymbolStruct = exchange.GetSymbols()

	//fmt.Println(symbols.Symbols)
	for k, v := range symbols.Symbols.(map[string]interface{}) {
		fmt.Println(k, "-", v)
	}
} */