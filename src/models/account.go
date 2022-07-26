package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID 				uint	`json:"id"`
	AccountType		string 	`json:"account_type" gorm:"not null"`
	AccountStatus	string	`json:"account_status" gorm:"not null"`
    Email 			string	`json:"email" gorm:"unique; not null"`
    Username 		string	`json:"username" gorm:"unique; not null"`
    Password 		[]byte	`json:"-" gorm:"not null"`
    ImagePath 		string	`json:"image_path" gorm:"not null"`
    Balance 		uint	`json:"balance" gorm:"not null"`
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