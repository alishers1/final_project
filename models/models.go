package models

import (
	"gorm.io/gorm"
)

type User struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number" gorm:"unique"`
	Email       string `json:"email" gorm:"unique"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Password    string `json:"-"`
	IsActive    bool   `json:"-"`
	gorm.Model
}

type Client struct {
	FullName    string   `json:"full_name"`
	TIN         string   `json:"tin" gorm:"unique"`
	PhoneNumber string   `json:"phone_number" gorm:"unique"`
	Email       string   `json:"email" gorm:"unique"`
	Products    []string `json:"products" gorm:"type:text[]"`
	Age         int      `json:"age"`
	Gender      string   `json:"gender"`
	UserID      uint     `json:"user_id" gorm:"foreignKey:ID"`
	Password    string   `json:"-"`
	IsActive    bool     `json:"-"`
	gorm.Model
}
