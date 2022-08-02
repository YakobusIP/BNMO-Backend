package models

import (
	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	RequestType		string	`json:"request_type" gorm:"not null"`
	Status         	string 	`json:"status" gorm:"not null"`
	Amount         	uint	`json:"amount" gorm:"not null"`
	Currency		string	`json:"currency" gorm:"not null"`
	ConvertedAmount	uint	`json:"converted_amount" gorm:"not null"`
	AccountID    	uint	`json:"account_id"`
	Account 		Account	`json:"account" gorm:"foreignKey:AccountID;references:ID"`
}