package helper

import (
	"os"

	"github.com/kelvins/geocoder"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Decrypt(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Geocode(city string) (float64, float64, error) {
	geocoder.ApiKey = os.Getenv("API_KEY")

	var lng, lat float64
	var address geocoder.Address
	address.City = city

	location, err := geocoder.Geocoding(address)
	if err != nil {
		return 0, 0, err
	} else {
		lng = location.Latitude
		lat = location.Longitude
	}
	return lng, lat, nil
}
