package user

import (
	"errors"
	"regexp"
	"shopping-api/internal/dto"
	"shopping-api/internal/factory"
	"shopping-api/internal/model"
	"shopping-api/internal/repository"
	"shopping-api/pkg/helper"
	"strings"

	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateUser(user model.Users) (err error)
	GetUser(id int) (*model.Users, error)
	DeleteUser(id int) error
	LoginUser(user dto.UserLogin) (interface{}, error)
}

type service struct {
	UserRepository repository.User
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) CreateUser(user model.Users) (err error) {
	spaceEmpty := strings.TrimSpace(user.Name)
	newPass, _ := helper.Encrypt(user.Password)
	user.Password = newPass

	if user.Name == "" && user.Email == "" && user.Password == "" && user.Phone_Number == "" {
		return errors.New("please fill user name and email")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	if len(user.Password) < 5 {
		return errors.New("password is too short")
	}

	if spaceEmpty == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Phone_Number == "" {
		return errors.New("phone number is required")
	}
	pattern := `^\w+@\w+\.\w+$`
	matched, _ := regexp.Match(pattern, []byte(user.Email))
	if !matched {
		return errors.New("format email is invalid")
	}
	createdUser, err := s.UserRepository.CreateUser(&user)
	if err != nil {
		logrus.Error(err)
		return err
	}
	cart := model.Cart{
		UsersID: createdUser.ID,
	}
	_, err = s.UserRepository.CreateCartUser(&cart)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (s *service) GetUser(id int) (*model.Users, error) {
	user, err := s.UserRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user.(*model.Users), nil
}

func (s *service) DeleteUser(id int) error {
	_, err := s.UserRepository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) LoginUser(user dto.UserLogin) (interface{}, error) {
	users, err := s.UserRepository.LoginUser(user)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return users, nil
}
