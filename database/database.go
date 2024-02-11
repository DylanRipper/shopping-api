package database

import (
	"fmt"
	"shopping-api/internal/model"
	"shopping-api/pkg/util"
	"sync"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	listener *pq.Listener
	once     sync.Once
)

func CreateConnection() {
	conf := dbConfig{
		User: util.GetEnv("PGSQL_USER", "open-api-xendit"),
		Pass: util.GetEnv("PGSQL_PASS", "komerce2023"),
		Host: util.GetEnv("PGSQL_HOST", "localhost"),
		Port: util.GetEnv("PGSQL_PORT", "5433"),
		Name: util.GetEnv("PGSQL_DB_NAME", "open-api-xendit"),
	}

	pgsql := pgsqlConfig{dbConfig: conf}

	once.Do(func() {
		pgsql.Connect()
	})

}

func GetConnection() *gorm.DB {
	if db == nil {
		CreateConnection()
		test, _ := db.DB()
		err := test.Ping()
		fmt.Println(err)

		err = listener.Ping()
		fmt.Println(err)
	}

	InitMigrate(db)

	return db
}

func InitMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&model.Category{})
	DB.AutoMigrate(&model.Subcategory{})
	DB.AutoMigrate(&model.Province{})
	DB.AutoMigrate(&model.City{})
	DB.AutoMigrate(&model.Users{})
	DB.AutoMigrate(&model.Photos{})
	DB.AutoMigrate(&model.Products{})
	DB.AutoMigrate(&model.Cart{})
	DB.AutoMigrate(&model.Transaction{})
	DB.AutoMigrate(&model.Booking{})
	DB.AutoMigrate(&model.Reviews{})
}
