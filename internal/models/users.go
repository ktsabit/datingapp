package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string `gorm:"uniqueIndex;not null" json:"email"`
	Password        string `gorm:"not null" json:"-"`
	Name            string `gorm:"not null" json:"name"`
	DailySwipeLimit int    `gorm:"not null" json:"daily_swipe_limit"`
	IsPremium       bool   `gorm:"default:false" json:"is_premium"`
	IsVerified      bool   `gorm:"default:false" json:"is_verified"`
}
