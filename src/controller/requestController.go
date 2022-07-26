package controller

import (
	"BNMO/database"
	"BNMO/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestAdd(c *gin.Context) {
	var request models.Request

	// Bind arriving json into a map
	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println("Unable to parse body into an account struct:" + err.Error())
		return
	}

	err = database.DATABASE.Create(&request).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request successfully added"})
}

func RequestSubtract(c *gin.Context) {
	var request models.Request

	// Bind arriving json into a map
	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println("Unable to parse body into an account struct:" + err.Error())
		return
	}

	err = database.DATABASE.Create(&request).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request successfully added"})
}
