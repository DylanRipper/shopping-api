package main

import (
	"shopping-api/database"
	"shopping-api/internal/factory"
	"shopping-api/internal/http"
	"shopping-api/pkg/util"
)

func main() {
	database.CreateConnection()

	f := factory.NewFactory()
	e := http.NewHttp(f)
	e.Logger.Fatal(e.Start(":" + util.GetEnv("APP_PORT", "8080")))
}
