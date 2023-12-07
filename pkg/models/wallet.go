package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model

	Balance float64 `json:"balance"`
}
