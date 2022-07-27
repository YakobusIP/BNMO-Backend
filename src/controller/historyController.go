package controller

import (
	"BNMO/database"
	"BNMO/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RequestHistory(c *gin.Context) {
	// Specify limitations
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 5
	offset := (page-1) * limit

	var total int64
	var getRequests []models.Request

	// Pull data from the requests table inside the database
	// Pull only based on the number of offsets and limits specified
	database.DATABASE.Offset(offset).Limit(limit).Find(&getRequests)
	database.DATABASE.Model(&models.Request{}).Count(&total)

	// Return data to frontend
	c.JSON(http.StatusOK, gin.H{
		"data": getRequests,
		"metadata": gin.H{
			"total": total,
			"page": page,
			"last_page": math.Ceil(float64(total)/float64(limit)),
		},
	})
}

func TransferHistory(c *gin.Context) {
	// Specify limitations
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 5
	offset := (page-1) * limit

	var total int64
	var getTransfers []models.Transfer

	// Pull data from the requests table inside the database
	// Pull only based on the number of offsets and limits specified
	database.DATABASE.Offset(offset).Limit(limit).Find(&getTransfers)
	database.DATABASE.Model(&models.Request{}).Count(&total)

	// Return data to frontend
	c.JSON(http.StatusOK, gin.H{
		"data": getTransfers,
		"metadata": gin.H{
			"total": total,
			"page": page,
			"last_page": math.Ceil(float64(total)/float64(limit)),
		},
	})
}