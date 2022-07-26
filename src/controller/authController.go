package controller

import (
	"BNMO/database"
	"BNMO/models"
	"BNMO/utilities"
	"fmt"
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
	
}

type Claims struct {
	jwt.StandardClaims
} 