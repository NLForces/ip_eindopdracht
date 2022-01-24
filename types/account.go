package types

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name      string
	Code      string
	Maxcredit float64
	Pincode   string
	CreatedAt time.Time
}
