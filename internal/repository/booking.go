package repository

import "gorm.io/gorm"

type Booking interface {
	GetProductID(id int) (int, error)
	GetBookingOwner(id int) (int, error)
	GetBookingStatus(id int) (string, error)
}

type booking struct {
	DB *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *booking {
	return &booking{
		DB: db,
	}
}

func (b *booking) GetProductID(id int) (int, error) {
	var productID int
	tx := b.DB.Raw("SELECT products_id FROM bookings WHERE id = ?", id).Scan(&productID)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return 0, tx.Error
	}
	return productID, nil
}

func (b *booking) GetBookingOwner(id int) (int, error) {
	var userID int
	tx := b.DB.Raw("SELECT carts.users_id FROM bookings JOIN carts ON bookings.cart_id = carts.id WHERE bookings.id = ?", id).Scan(&userID)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return 0, tx.Error
	}
	return userID, nil
}

func (b *booking) GetBookingStatus(id int) (string, error) {
	var statusPayment string
	tx := b.DB.Raw("SELECT status_payment FROM bookings WHERE id = ?", id).Scan(&statusPayment)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return "", tx.Error
	}
	return statusPayment, nil
}
