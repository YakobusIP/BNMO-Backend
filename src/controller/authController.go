package controller

import (
	"BNMO/database"
	"BNMO/models"
	"BNMO/utilities"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+\-]+@[a-z0-9. %+\-]+\.[a-z0-9. %+\-]`)
	return Re.MatchString(email)
}

func RegisterAccount(c *gin.Context) {
	var data map[string]interface{}
	var accountData models.Account

	// Bind arriving json into a map
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("Unable to parse body into an account struct:" + err.Error())
		return
	}

	// Check if the length of password is less than 8 characters
	if len(data["password"].(string)) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password must be 8 characters or more"})
		return
	}

	// Validate the email
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email address"})
		return
	}

	// Check if email already exist within the database
	database.DATABASE.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&accountData)
	if accountData.ID !=0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	// Check if username already exist within the database
	database.DATABASE.Where("username=?", data["username"].(string)).First(&accountData)
	if accountData.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username already exists"})
		return
	}
	
	// Map json into another variable
	account := models.Account{
		AccountType: data["account_type"].(string),
		AccountStatus: data["account_status"].(string),
		FirstName: data["first_name"].(string),
		LastName: data["last_name"].(string),
		Email: strings.TrimSpace(data["email"].(string)),
		Username: data["username"].(string),
		ImagePath: data["image_path"].(string),
		Balance: uint(data["balance"].(float64)),
	}
	// Hash password using bcrypt
	account.SetPassword(data["password"].(string))

	insert := database.DATABASE.Create(&account)
	if insert.Error != nil {
		fmt.Println("Failed to insert into database:" + insert.Error.Error())
	}

	c.JSON(http.StatusOK, gin.H{"account": account,
		"message": "Account successfully registered"})
}

func LoginAccount(c *gin.Context) {
	var data map[string]string
	var account models.Account

	// Bind arriving json into a map
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("Unable to parse body into an account struct:" + err.Error())
		return
	}

	// Check if email exists inside the database
	database.DATABASE.Where("email=?", data["email"]).First(&account)
	if account.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email does not exist"})
		return
	}

	// Check password validity
	err = account.ComparePassword(data["password"]) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect password"})
		return
	}

	if account.AccountStatus == "accepted" {	
		// Authenticate user
		token, err := utilities.GenerateJWT(strconv.Itoa(int(account.ID)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		// Set cookies
		c.SetCookie("jwt", token, int(time.Now().Add(time.Hour * 24).Unix()), "", "", true, true)
		c.JSON(http.StatusOK, gin.H{"account": account,
			"message": "Login successful"})
	} else if account.AccountStatus == "pending" {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Account isn't verified. Please wait for validation"})
		return
	} else if account.AccountStatus == "rejected" {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Account is rejected. Please contact our support"})
		return
	} 
}

type Claims struct {
	jwt.StandardClaims
} 

func DisplayPendingAccount(c *gin.Context) {
	// Specify limitations
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 5
	offset := (page-1) * limit

	var total int64
	var getAccounts []models.Account

	// Pull data from the requests table inside the database
	// Pull only based on the number of offsets and limits specified
	database.DATABASE.Preload("Accounts").Offset(offset).Limit(limit).Where("account_status=?", "pending").Find(&getAccounts)
	database.DATABASE.Model(&models.Account{}).Where("account_status=?", "pending").Count(&total)

	// Return data to frontend
	c.JSON(http.StatusOK, gin.H{
		"data": getAccounts,
		"metadata": gin.H{
			"total": total,
			"page": page,
			"last_page": math.Ceil(float64(int(total)/limit)),
		},
	})
}

func ValidateAccount(c *gin.Context) {
	var data map[string]interface{}
	var account models.Account

	// Bind arriving json into a map
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("Unable to parse body into a validate_request struct:" + err.Error())
		return
	}

	if data["validation"] == "accepted" {
		database.DATABASE.First(&account, uint(data["id"].(float64))).Update("account_status", "accepted")
	} else if data["validation"] == "rejected" {
		database.DATABASE.First(&account, uint(data["id"].(float64))).Update("account_status", "rejected")
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account validation successful"})
}