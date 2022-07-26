package routes

import (
	"BNMO/controller"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func MapUrls(Router *gin.Engine) {
	Router.POST("/api/register", controller.RegisterAccount)
	Router.POST("/api/login", controller.LoginAccount)
	
	Router.Use(middleware.IsAuthenticate)
	Router.POST("/api/customerrequest/add", controller.RequestAdd)
	Router.POST("/api/customerrequest/subtract", controller.RequestSubtract)
}