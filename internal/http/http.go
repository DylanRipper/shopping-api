package http

import (
	"shopping-api/internal/app/product"
	"shopping-api/internal/app/user"
	"shopping-api/internal/factory"
	middlewares "shopping-api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func NewHttp(f *factory.Factory) *echo.Echo {
	e := middlewares.CORSMiddleware()
	v1 := e.Group("/api/v1")
	user.NewHandler(f).UserRouter(v1.Group("/general-user"))
	user.NewHandler(f).JwtUserRouter(v1.Group("/jwt-user"))
	product.NewHandler(f).ProductRouter(v1.Group("/general-product"))
	product.NewHandler(f).JwtProductRouter(v1.Group("/jwt-product"))

	return e
}
