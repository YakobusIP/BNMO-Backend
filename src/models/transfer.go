package models

import (
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	ID				uint		`json:"id"`
	AccountID 		string		`json:"account_id"`
    Destination 	string		`json:"destination"`
    Amount 			uint		`json:"amount" gorm:"not null"`
	Currency		string		`json:"currency" gorm:"not null"`
	ConvertedAmount	uint		`json:"converted_amount" gorm:"not null"`
	Status			string		`json:"status" gorm:"not null"`
	Account 		Account		`json:"account" gorm:"foreignKey:AccountID"`
}