package models

import "gorm.io/gorm"

type Swipe struct {
	gorm.Model
	SwiperID uint `gorm:"index"`
	SwipedID uint `gorm:"index"`
}
