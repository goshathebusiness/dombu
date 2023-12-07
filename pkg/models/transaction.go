package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	BalanceID   uint    `json:"balance_id"`
	Topic       string  `json:"topic"`
	MoneyChange float64 `json:"money_change"`
}
