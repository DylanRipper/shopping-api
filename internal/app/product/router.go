package product

import (
	"shopping-api/pkg/constant"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *handler) ProductRouter(r *echo.Group) {
	r.GET("/all-products", h.GetAllProduct)
	r.GET("/products", h.GetProductsByUserID)
	r.GET("/products/:id", h.GetProductByID)
	r.GET("/products/subcategory/:id", h.GetProductsBySubcategoryID)
	r.POST("/search", h.GetProductsByName)

}

func (h *handler) JwtProductRouter(r *echo.Group) {
	r.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	r.POST("/products", h.CreateProduct)
	r.GET("/products", h.GetProductsByUserID)
	r.DELETE("/products/:id", h.DeleteProductByID)
}
