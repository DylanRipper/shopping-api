package factory

import (
	"shopping-api/database"
	"shopping-api/internal/repository"
)

type Factory struct {
	UserRepository    repository.User
	ProductRepository repository.Product
	CityRepository    repository.City
	ReviewRepository  repository.Review
	PhotoRepository   repository.Photo
	BookingRepository repository.Booking
}

func NewFactory() *Factory {
	db := database.GetConnection()

	return &Factory{
		UserRepository:    repository.NewUserRepository(db),
		ProductRepository: repository.NewProductRepository(db),
		CityRepository:    repository.NewCityRepository(db),
		ReviewRepository:  repository.NewReviewRepository(db),
		PhotoRepository:   repository.NewPhotoRepository(db),
		BookingRepository: repository.NewBookingRepository(db),
	}

}
