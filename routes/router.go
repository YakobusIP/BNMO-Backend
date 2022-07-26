package routes

import (
	"BNMO/controller"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func MapUrls(Router *gin.Engine) {
	// Register account
	Router.POST("/api/register", controller.RegisterAccount)
	// Upload image
	Router.POST("/api/upload-image", controller.Upload)
	// Login account
	Router.POST("/api/login", controller.LoginAccount)
	// Logout accounts
	Router.POST("/api/logout", controller.LogoutAccount)
	// Get image
	Router.Static("/api/uploads", "images")
	
	// Authorized side (both customer and admin)
	Router.Use(middleware.IsAuthenticate)
	/* CUSTOMER SIDE */
	// Show user profile
	Router.GET("/api/profile/:id", controller.ShowProfile)
	// Request add balance
	Router.POST("/api/customerrequest/add", controller.CustomerRequest)
	// Request subtract balance
	Router.POST("/api/customerrequest/subtract", controller.CustomerRequest)
	// Get data on accounts for transfer purposes
	Router.GET("/api/displayaccounts", controller.DisplayAccounts)
	// Transfer from source account to destination account
	Router.POST("/api/transfer", controller.Transfer)
	// Get request history
	Router.GET("/api/requesthistory", controller.RequestHistory)
	// Get transfer history
	Router.GET("/api/transferhistory", controller.TransferHistory)
	// Update image
	Router.POST("/api/updateimage", controller.UpdateImage)

	/* ADMIN SIDE */
	// Display requests for admin
	Router.GET("/api/displayrequest", controller.DisplayRequests)
	// Validate selected requests for admin
	Router.POST("/api/validaterequest", controller.ValidateRequests)
	// Display pending accounts for admin
	Router.GET("/api/displaypending", controller.DisplayPendingAccount)
	// Validate selected accounts for admin
	Router.POST("/api/validateaccount", controller.ValidateAccount)
	// Display all customer data
	Router.GET("/api/customerdata", controller.SendAllCustomerData)
}