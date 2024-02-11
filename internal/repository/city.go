package repository

import (
	"shopping-api/internal/model"

	"gorm.io/gorm"
)

type City interface {
	GetOneCity(selectField, query string, args ...any) (getCity *model.City, err error)
	GetManyCity(selectField, query string, args ...any) (getCities []model.City, err error)
}

type city struct {
	DB *gorm.DB
}

func NewCityRepository(db *gorm.DB) *city {
	return &city{
		DB: db,
	}
}

func (c *city) GetOneCity(selectField, query string, args ...any) (getCity *model.City, err error) {
	err = c.DB.Model(&model.City{}).Select(selectField).Where(query, args...).Last(&getCity).Error
	if err != nil {
		return nil, err
	}

	return getCity, nil
}

func (c *city) GetManyCity(selectField, query string, args ...any) (getCities []model.City, err error) {
	err = c.DB.Model(&model.City{}).Select(selectField).Where(query, args...).Find(&getCities).Error
	if err != nil {
		return nil, err
	}

	return getCities, nil
}
