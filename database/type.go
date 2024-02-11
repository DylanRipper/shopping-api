package database

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	dbConfig struct {
		Host string
		User string
		Pass string
		Port string
		Name string
	}

	pgsqlConfig struct {
		dbConfig
	}
)

func (p pgsqlConfig) Connect() {
	conninfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=public",
		p.Host,
		p.Port,
		p.User,
		p.Pass,
		p.Name)

	fmt.Println(conninfo)
	pgx, err := sql.Open("pgx", conninfo)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn:                 pgx,
		WithoutReturning:     false,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
		Logger:                   logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		logrus.Error(err)
		panic(err)
	} else {
		logrus.Println("Connected to database PostgreSQL")
	}
}
