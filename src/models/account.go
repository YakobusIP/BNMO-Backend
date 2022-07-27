package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID 				uint	`json:"id"`
	AccountType		string 	`json:"account_type" gorm:"not null"`
	AccountStatus	string	`json:"account_status" gorm:"not null"`
	FirstName		string	`json:"first_name" gorm:"not null"`
	LastName		string	`json:"last_name" gorm:"not null"`
    Email 			string	`json:"email" gorm:"unique; not null"`
    Username 		string	`json:"username" gorm:"unique; not null"`
    Password 		[]byte	`json:"-" gorm:"not null"`
    ImagePath 		string	`json:"image_path" gorm:"not null"`
    Balance 		uint	`json:"balance" gorm:"not null"`
}

type TransferAccount struct {
	ID				uint	`json:"id"`
	FirstName		string	`json:"first_name"`
	LastName		string	`json:"last_name"`
	Username		string	`json:"username"`
}

type CustomerData struct {
	ID				uint		`json:"id"`
	FirstName		string		`json:"first_name"`
	LastName		string		`json:"last_name"`
	Username		string		`json:"username"`
	Email 			string		`json:"email"`
	ImagePath 		string		`json:"image_path"`
    Balance 		uint		`json:"balance"`
	CreatedAt		time.Time	`json:"created_at"`
}

// Function to hash password using bcrypt with salt
func (account *Account) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	account.Password = hashedPassword
}

// Function to compare user inputted password with the one inside the database
func (account *Account) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(account.Password, []byte(password))
}