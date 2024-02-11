package repository

import (
	"shopping-api/internal/dto"
	middlewares "shopping-api/internal/middleware"
	"shopping-api/internal/model"
	"shopping-api/pkg/helper"

	"gorm.io/gorm"
)

type User interface {
	CreateUser(user *model.Users) (*model.Users, error)
	CreateCartUser(cart *model.Cart) (interface{}, error)
	GetUser(id int) (interface{}, error)
	GetUserByEmail(loginuser model.Users) (*model.Users, error)
	UpdateUser(id int, user *model.Users) (interface{}, error)
	DeleteUser(id int) (interface{}, error)
	LoginUser(UserLogin dto.UserLogin) (interface{}, error)
}

type user struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{
		DB: db,
	}
}

// function database untuk menambahkan user baru (registrasi)
func (u *user) CreateUser(user *model.Users) (*model.Users, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// function database untuk membuat cart user
func (u *user) CreateCartUser(cart *model.Cart) (interface{}, error) {
	if err := u.DB.Create(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

// function database untuk menampilkan user by id
func (u *user) GetUser(id int) (interface{}, error) {
	var user model.Users
	var result dto.Get_User
	err := u.DB.Model(user).Find(&result, id)
	rows_affected := err.RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return result, nil
}

func (u *user) GetUserByEmail(loginuser model.Users) (*model.Users, error) {
	var user model.Users
	tx := u.DB.Where("email=?", loginuser.Email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	checkpass := helper.Decrypt(loginuser.Password, user.Password)
	if !checkpass {
		return nil, nil
	}
	return &user, nil
}

// function database untuk memperbarui data user by id
func (u *user) UpdateUser(id int, user *model.Users) (interface{}, error) {
	if err := u.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}
	u.DB.First(&user, id)
	return user, nil
}

// function database untuk menghapus data user by id
func (u *user) DeleteUser(id int) (interface{}, error) {
	var user model.Users
	if err := u.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// function database untuk melakukan login
func (u *user) LoginUser(UserLogin dto.UserLogin) (interface{}, error) {
	var result dto.Get_User
	var user model.Users
	var err error
	if err = u.DB.Where("email = ?", UserLogin.Email).Find(&user).Error; err != nil {
		return nil, err
	}

	check := helper.Decrypt(user.Password, UserLogin.Password)
	if !check {
		return 0, nil
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := u.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	u.DB.Model(user).Find(&result, user)
	return result, nil
}
