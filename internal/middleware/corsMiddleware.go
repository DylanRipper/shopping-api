package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func CORSMiddleware() *echo.Echo {
	e := echo.New()
	e.Use(m.CORSWithConfig(m.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	return e
}
