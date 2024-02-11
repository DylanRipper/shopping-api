package util

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {
	a, _ := godotenv.Read()

	var (
		val     string
		isExist bool
	)

	val, isExist = a[key]
	if !isExist {
		val = os.Getenv(key)
		if val == "" {
			val = fallback
		}
	}

	return val

}
