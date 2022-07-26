package models

import (
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	ID				uint		`json:"id"`
	AccountID 		uint		`json:"account_id"`
    Destination 	uint		`json:"destination"`
    Amount 			uint		`json:"amount" gorm:"not null"`
	Currency		string		`json:"currency" gorm:"not null"`
	Account 		Account		`json:"account" gorm:"foreignKey:AccountID"`
}