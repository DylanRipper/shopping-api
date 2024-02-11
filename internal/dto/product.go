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
	}

	// struct get product
	GetAllProduct struct {
		ID               uint   `json:"id" form:"id"`
		UsersID          uint   `json:"users_id" form:"users_id"`
		Name             string `json:"name" form:"name"`
		Subcategory_Name string `json:"subcategory_name" form:"subcategory_name"`
		SubcategoryID    int    `json:"subcategory_id" form:"subcategory_id"`
		CityID           int    `json:"city_id" form:"city_id"`
		City_Name        string `json:"city_name" form:"city_name"`
		Price            int    `json:"price" form:"price"`
		Description      string `json:"description" form:"description"`
		Stock            int    `json:"stock" form:"stock"`
		Url              string `json:"url" form:"url"`
	}

	// struct get product
	GetProduct struct {
		ID               uint      `json:"id" form:"id"`
		UsersID          uint      `json:"users_id" form:"users_id"`
		CreatedAt        time.Time `json:"created_at" form:"created_at"`
		PhoneNumber      string    `json:"phone_number" form:"phone_number"`
		Name             string    `json:"name" form:"name"`
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

	SearchName struct {
		Name string `query:"name" json:"name" form:"name"`
	}
)
