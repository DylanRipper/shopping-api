package user

import (
	"shopping-api/pkg/constant"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *handler) UserRouter(r *echo.Group) {
	r.POST("/signup", h.CreateUser)
	r.POST("/signin", h.LoginUser)
}

func (h *handler) JwtUserRouter(r *echo.Group) {
	r.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	r.GET("/users", h.GetUser)
	r.DELETE("/users", h.DeleteUser)

}
