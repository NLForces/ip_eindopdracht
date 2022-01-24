package types

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	AccountId   uint
	Amount      float64
	Type        string `gorm:"type:enum('debet', 'credit');default:'debet'"`
	Description string
}
