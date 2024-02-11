package model

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ProductsID     uint      `json:"product_id" form:"product_id"`
	CartID         uint      `json:"cart_id" form:"cart_id"`
	UsersID        uint      `json:"user_id" form:"user_id"`
	TransactionID  *uint     `json:"transaction_id" form:"transaction_id"`
	Time_In        time.Time `gorm:"type:datetime;not null" json:"time_in" form:"time_in"`
	Time_Out       time.Time `gorm:"type:datetime;not null" json:"time_out" form:"time_out"`
	Total_Day      int       `json:"total_day" form:"total_day"`
	Qty            int       `gorm:"not null" json:"qty" form:"qty"`
	Total          int       `json:"total" form:"total"`
	Status_Payment string    `gorm:"type:enum('waiting','succes');default:'waiting';not null" json:"status_payment" form:"status_payment"`
	Review         Reviews   `gorm:"foreignKey:BookingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
