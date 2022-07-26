package models

import (
	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	ID				uint	`json:"id"`
	AccountID    	string	`json:"account_id"`
	RequestType		string	`json:"request_type" gorm:"not null"`
	Status         	string 	`json:"status" gorm:"not null"`
	Amount         	uint	`json:"amount" gorm:"not null"`
	Currency		string	`json:"currency" gorm:"not null"`
	Account 		Account	`json:"account" gorm:"foreignKey:AccountID"`
}