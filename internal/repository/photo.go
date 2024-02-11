package repository

import (
	"shopping-api/internal/model"

	"gorm.io/gorm"
)

type Photo interface {
	InsertPhoto(photo *model.Photos) (interface{}, error)
	GetUrl(id uint) ([]string, error)
}

type photo struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photo {
	return &photo{
		DB: db,
	}
}

func (p *photo) InsertPhoto(photo *model.Photos) (interface{}, error) {
	if err := p.DB.Create(&photo).Error; err != nil {
		return nil, err
	}
	return photo, nil
}

// Fungsi untuk mendapatkan seluruh url photo product tertentu
func (p *photo) GetUrl(id uint) ([]string, error) {
	var url []string
	tx := p.DB.Table("photos").Select("photos.url").Where("photos.products_id = ?", id).Find(&url)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return url, nil
}
