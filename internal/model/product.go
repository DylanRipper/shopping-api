package model

import (
	"gorm.io/gorm"
)

// struct user
type Products struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(255)" json:"name" form:"name"`
	UsersID       uint      `json:"user_id" form:"user_id"`
	SubcategoryID int       `json:"subcategory_id" form:"subcategory_id"`
	CityID        int       `json:"city_id" form:"city_id"`
	Price         int       `gorm:"not null" json:"price" form:"price"`
	Description   string    `gorm:"type:text;not null" json:"description" form:"description"`
	Stock         int       `gorm:"type:int;default:1" json:"stock" form:"stock"`
	Longitude     float64   `gorm:"type:varchar(30);not null" json:"lon" form:"lon"`
	Latitude      float64   `gorm:"type:varchar(30);not null" json:"lat" form:"lat"`
	Rating        float64   `gorm:"type:decimal(2,1)" json:"rating" form:"rating"`
	Photos        []Photos  `gorm:"foreignKey:ProductsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Booking       []Booking `gorm:"foreignKey:ProductsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Reviews       []Reviews `gorm:"foreignKey:ProductsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
