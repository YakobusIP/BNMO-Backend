package controller

import (
	"BNMO/database"
	"BNMO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowProfile(c *gin.Context) {
	strid, _ := c.Params.Get("id")
	id, _ := strconv.Atoi(strid)
	var account models.Account
	database.DATABASE.Where("id=?", id).Find(&account)
	c.JSON(http.StatusOK, gin.H{
		"data": account,
	})
}

func SendAllCustomerData(c *gin.Context) {
	var account []models.Account
	var getCustomerData []models.CustomerData

	// Pull data from the requests table inside the database
	database.DATABASE.Where(map[string]string{"account_status": "accepted", "account_type": "customer"}).Select("id", "first_name", "last_name", "username", "email", "image_path", "balance", "created_at").Find(&account).Scan(&getCustomerData)

	// Return data to frontend
	c.JSON(http.StatusOK, gin.H{
		"data": getCustomerData,
	})
}