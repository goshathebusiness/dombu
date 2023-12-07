package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	WalletID    uint    `json:"wallet_id"`
	Topic       string  `json:"topic"`
	MoneyChange float64 `json:"money_change"`
}
