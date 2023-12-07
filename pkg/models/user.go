package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Credentials *Credentials `json:"credentials"`
	WalletID    uint         `json:"wallet_id"`
}

type Credentials struct {
	gorm.Model

	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
