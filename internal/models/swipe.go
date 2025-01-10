package models

import "gorm.io/gorm"

type Swipe struct {
	gorm.Model
	UserID         uint               `gorm:"foreignKey:ProfileID"`
	TargetID       uint               `gorm:"foreignKey:ProfileID"`
	SwipeDirection SwipeDirectionEnum `gorm:"type:varchar(10)"`
	IsMatch        bool               `gorm:"default:false"`
}

type SwipeDirectionEnum string

const (
	SwipeRight SwipeDirectionEnum = "right"
	SwipeLeft  SwipeDirectionEnum = "left"
)
