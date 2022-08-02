package models

import (
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
    Destination 	string		`json:"destination"`
    Amount 			uint		`json:"amount" gorm:"not null"`
	Currency		string		`json:"currency" gorm:"not null"`
	ConvertedAmount	uint		`json:"converted_amount" gorm:"not null"`
	Status			string		`json:"status" gorm:"not null"`
	AccountID 		uint		`json:"account_id"`
	Account 		Account		`json:"account" gorm:"foreignKey:AccountID;references:ID"`
}