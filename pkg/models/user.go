package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Credentials *Credentials
	WalletID    uint
}

type Credentials struct {
	gorm.Model

	Email    string
	Password string
}
