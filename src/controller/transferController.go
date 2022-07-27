package controller

import (
	"BNMO/database"
	"BNMO/models"
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayAccounts(c *gin.Context) {
	var account []models.Account
	var transferAccounts []models.TransferAccount
	database.DATABASE.Where(map[string]string{"account_type":"customer", "account_status": "accepted"}).Select("id", "first_name", "last_name", "username").Find(&account).Scan(&transferAccounts)

	c.JSON(http.StatusOK, gin.H{
		"data": transferAccounts,
	})
}

func Transfer(c *gin.Context) {
	var source models.Account
	var destination models.Account
	var transfer models.Transfer

	// Bind arriving json into a map
	err := c.BindJSON(&transfer)
	if err != nil {
		fmt.Println("Unable to parse body into a validate_request struct:" + err.Error())
		return
	}

	// Pull data from accounts table
	database.DATABASE.Create(&transfer)
	database.DATABASE.Find(&source, transfer.AccountID)
	database.DATABASE.Find(&destination, transfer.Destination)

	// Pull rates from redis
	_, conversionRates := getRatesFromRedis(transfer.Currency)
	convertedAmount := float64(transfer.Amount) / conversionRates

	// If balance is insufficient
	if source.Balance < uint(math.Floor(convertedAmount)) {
		database.DATABASE.Model(&transfer).Update("status", "failed")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Insufficient balance"})
		return
	}

	// Subtract balance from source
	// Add balance to destination
	newSourceBalance := source.Balance - uint(convertedAmount)
	newDestinationBalance := destination.Balance + uint(convertedAmount)

	// Update database values
	database.DATABASE.Find(&source, transfer.AccountID).Update("balance", newSourceBalance)
	database.DATABASE.Find(&destination, transfer.Destination).Update("balance", newDestinationBalance)
	database.DATABASE.Model(&transfer).Update("status", "success")
	database.DATABASE.Model(&transfer).Update("converted_amount", uint(math.Floor(convertedAmount)))
	c.JSON(http.StatusOK, gin.H{"message": "Transfer completed"})
}