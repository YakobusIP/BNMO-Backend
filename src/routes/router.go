package routes

import (
	"github.com/YakobusIP/BNMO-Backend.git/controller"
	"github.com/gin-gonic/gin"
)

func MapUrls(Router *gin.Engine) {
	Router.GET("/api/v1/customerrequest", controller.GetSymbolsFromRedis)
}