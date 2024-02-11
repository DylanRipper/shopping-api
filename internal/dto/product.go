package dto

import "time"

type (
	// struct body create product
	BodyCreateProducts struct {
		Name          string `gorm:"type:varchar(255)" json:"name" form:"name"`
		SubcategoryID int    `json:"subcategory_id" form:"subcategory_id"`
		CityID        int    `json:"city_id" form:"city_id"`
		Price         int    `gorm:"not null" json:"price" form:"price"`
		Description   string `gorm:"type:text;not null" json:"description" form:"description"`
		Stock         int    `gorm:"type:int;default:1" json:"stock" form:"stock"`
		Guarantee     []int  `json:"guarantee" form:"guarantee"`
	}

	// struct get product
	GetAllProduct struct {
		ID               uint
		UsersID          uint
		Name             string
		Subcategory_Name string
		SubcategoryID    int
		CityID           int
		City_Name        string
		Price            int
		Description      string
		Stock            int
		Url              string
	}

	// struct get product
	GetProduct struct {
		ID               uint
		UsersID          uint
		CreatedAt        time.Time
		Nama             string
		Phone_Number     string
		Name             string
		SubcategoryID    int
		Subcategory_Name string
		CityID           int
		City_Name        string
		Price            int
		Description      string
		Stock            int
		Longitude        float64
		Latitude         float64
		Url              []string
		Guarantee        []string
	}
)
