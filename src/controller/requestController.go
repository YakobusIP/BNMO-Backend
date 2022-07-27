package controller

import (
	"BNMO/database"
	"BNMO/models"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Customer add and subtract requests
func CustomerRequest(c *gin.Context) {
	var request models.Request

	// Bind arriving json into a map
	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println("Unable to parse body into a request struct:" + err.Error())
		return
	}

	err = database.DATABASE.Create(&request).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request successfully added"})
}

// Admin display requests
func DisplayRequests(c *gin.Context) {
	// Specify limitations
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 5
	offset := (page-1) * limit

	var total int64
	var getRequests []models.Request

	// Pull data from the requests table inside the database
	// Pull only based on the number of offsets and limits specified
	database.DATABASE.Preload("Requests").Offset(offset).Limit(limit).Where("status=?", "pending").Find(&getRequests)
	database.DATABASE.Model(&models.Request{}).Where("status=?", "pending").Count(&total)

	// Return data to frontend
	c.JSON(http.StatusOK, gin.H{
		"data": getRequests,
		"metadata": gin.H{
			"total": total,
			"page": page,
			"last_page": math.Ceil(float64(int(total)/limit)),
		},
	})
}

// Admin accept or reject requests
func ValidateRequests(c *gin.Context) {
	var data map[string]interface{}
	var account models.Account
	var request models.Request

	// Bind arriving json into a map
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("Unable to parse body into a validate_request struct:" + err.Error())
		return
	}

	// If status is accepted, start procedures
	if data["validation"] == "accepted" {
		var convertedAmount float64
		// Check statements
		// Pull data from request and account tables
		database.DATABASE.First(&request, uint(data["id"].(float64)))
		database.DATABASE.First(&account, request.AccountID)
		source, conversionRates := GetRatesFromRedis(request.Currency)
		fmt.Println(source)

		// Request type: add
		if request.RequestType == "add" {
			convertedAmount = float64(request.Amount) / conversionRates
			newBalance := account.Balance + uint(math.Floor(convertedAmount))
			database.DATABASE.First(&account, request.AccountID).Update("balance", newBalance)
		}

		// Request type: subtract
		if request.RequestType == "subtract" {
			// If balance is insufficient, reject the request
			if account.Balance < request.Amount {
				database.DATABASE.First(&request, data["id"]).Update("status", "rejected")
				c.JSON(http.StatusBadRequest, gin.H{"message": "Insufficient balance"})
				return
			} else {
				convertedAmount = float64(request.Amount) / conversionRates
				newBalance := account.Balance - uint(math.Floor(convertedAmount))
				database.DATABASE.First(&account, request.AccountID).Update("balance", newBalance)
			}
		}

		// Update value inside request table
		database.DATABASE.First(&request, uint(data["id"].(float64))).Update("status", data["validation"].(string))
		database.DATABASE.First(&request, uint(data["id"].(float64))).Update("converted_amount", uint(math.Floor(convertedAmount)))
		c.JSON(http.StatusOK, gin.H{"message": "Successfully accepted"})
		return
		
	} else if data["validation"] == "rejected" {
		// Update value inside request table
		database.DATABASE.First(&request, uint(data["id"].(float64))).Update("status", data["validation"].(string))
		c.JSON(http.StatusOK, gin.H{"message": "Successfully rejected"})
		return
	}
	
}