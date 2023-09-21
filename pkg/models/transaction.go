package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	Topic       string
	MoneyChange float64
}
