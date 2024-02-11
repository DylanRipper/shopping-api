package repository

import (
	"shopping-api/internal/model"

	"gorm.io/gorm"
)

type Review interface {
	AddReviews(review *model.Reviews) (interface{}, error)
	AddRatingToProduct(id int)
}

type review struct {
	DB *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *review {
	return &review{
		DB: db,
	}
}

func (r *review) AddReviews(review *model.Reviews) (interface{}, error) {
	if err := r.DB.Create(&review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (r *review) AddRatingToProduct(id int) {
	r.DB.Exec("UPDATE products SET rating = (SELECT AVG(rating) FROM reviews WHERE products_id = ?) WHERE id = ?", id, id)
}
