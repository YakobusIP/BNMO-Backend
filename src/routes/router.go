package routes

import (
	"BNMO/controller"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func MapUrls(Router *gin.Engine) {
	Router.POST("/api/register", controller.RegisterAccount)
	Router.POST("/api/login", controller.LoginAccount)
	Router.GET("/api/displayrequest", controller.DisplayRequests)
	Router.POST("/api/validaterequest", controller.ValidateRequests)
	Router.GET("/api/displaypending", controller.DisplayPendingAccount)
	Router.POST("/api/validateaccount", controller.ValidateAccount)
	
	// Customer side
	Router.Use(middleware.IsAuthenticate)
	Router.POST("/api/customerrequest/add", controller.CustomerRequest)
	Router.POST("/api/customerrequest/subtract", controller.CustomerRequest)
	Router.GET("/api/displayaccounts", controller.DisplayAccounts)
	Router.POST("/api/transfer", controller.Transfer)
	Router.GET("/api/requesthistory", controller.RequestHistory)
	Router.GET("/api/transferhistory", controller.TransferHistory)
}