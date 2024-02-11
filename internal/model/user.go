package model

import "gorm.io/gorm"

// struct user
type Users struct {
	gorm.Model
	Name         string     `gorm:"type:varchar(255)" json:"name" form:"name"`
	Email        string     `gorm:"type:varchar(100);unique;not null" json:"email" form:"email"`
	Password     string     `gorm:"type:text;not null" json:"password" form:"password"`
	Phone_Number string     `gorm:"type:varchar(100);unique;not null" json:"phone" form:"phone"`
	Token        string     `gorm:"type:text" json:"token" form:"token"`
	Products     []Products `gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cart         Cart       `gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
