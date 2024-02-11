package dto

type (
	UserLogin struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	Get_User struct {
		ID          uint   `json:"id" form:"id"`
		Name        string `json:"name" form:"name"`
		Email       string `json:"email" form:"email"`
		PhoneNumber string `json:"phone_number" form:"phone_number"`
		Token       string `json:"token" form:"token"`
	}
)
