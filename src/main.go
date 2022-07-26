package main

import (
	"BNMO/database"
	"BNMO/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
)

func main() {
	database.Initialize()
	Router.Use(cors.Default())
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