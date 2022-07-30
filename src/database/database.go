package database

import (
	"BNMO/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DATABASE *gorm.DB
)

func Initialize() {
	dsn := "root:YIP190602@tcp(localhost)/bnmo?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to database using gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error opening database connection")
	} else {
		fmt.Println("Connected successfully")
	}
	db.AutoMigrate(
		&models.Account{},
		&models.Request{},
		&models.Transfer{})

	DATABASE = db
	seed(DATABASE)
}

func seed(db *gorm.DB) {
	accounts := models.Account{
		AccountType: "admin",
		AccountStatus: "accepted",
		FirstName: "Admin",
		LastName: "Admin",
		Email: "admin@gmail.com", 
		Username: "admin",
		Password: []byte("admin"),
		ImagePath: "./images/Admin.png",
		Balance: 0,
	}

	db.Create(&accounts)
}