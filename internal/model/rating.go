package model

import (
	"gorm.io/gorm"
)

type Reviews struct {
	gorm.Model
	BookingID  uint   `gorm:"unique" json:"booking_id" form:"booking_id"`
	Rating     int    `gorm:"type:int;not null" json:"rating" form:"rating"`
	Comment    string `gorm:"type:text;not null" json:"comment" form:"comment"`
	ProductsID uint   `json:"product_id" form:"product_id"`
}
