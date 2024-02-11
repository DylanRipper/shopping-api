package product

import (
	"net/http"
	"shopping-api/internal/dto"
	"shopping-api/internal/factory"
	middlewares "shopping-api/internal/middleware"
	"shopping-api/internal/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) CreateProduct(c echo.Context) error {
	var body dto.BodyCreateProducts
	userID := middlewares.ExtractTokenUserId(c)

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.StatusBadRequest("bad request.", err))

	}

	code, err := h.service.CreateProduct(body, userID)
	if err != nil {
		if code == http.StatusBadRequest {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest("bad request.", err))

		} else if code == http.StatusInternalServerError {
			return c.JSON(http.StatusInternalServerError, response.StatusInternalServerError("internal server error.", err))

		}
	}

	return c.JSON(http.StatusOK, response.StatusOK("successfully created product."))
}

func (h *handler) GetAllProduct(c echo.Context) error {
	product, err := h.service.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusInternalServerError("internal server error.", err))
	}

	return c.JSON(http.StatusOK, response.StatusOKWithData("success.", product))
}

func (h *handler) GetProductByID(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}

	product, err := h.service.GetProductByID(productId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusInternalServerError("internal server error.", err))
	}

	return c.JSON(http.StatusOK, response.StatusOKWithData("success.", product))
}

func (h *handler) GetProductsBySubcategoryID(c echo.Context) error {
	subcategoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	product, err := h.service.GetProductsBySubcategoryID(subcategoryId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusInternalServerError("internal server error.", err))
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, response.StatusNotFound("product not found", nil))
	}
	return c.JSON(http.StatusOK, response.StatusOKWithData("success.", product))
}

func (h *handler) GetProductsByUserID(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	product, err := h.service.GetProductsByUserID(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusInternalServerError("internal server error.", err))
	}

	return c.JSON(http.StatusOK, response.StatusOKWithData("success.", product))
}

func (h *handler) DeleteProductByID(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}

	userID := middlewares.ExtractTokenUserId(c)

	err = h.service.DeleteProductByID(productID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusInternalServerError("internal server error.", err))
	}

	return c.JSON(http.StatusOK, response.StatusOK("success delete product."))
}

// untuk mendapatkan product berdasarkan nama product
func (h *handler) GetProductsByName(c echo.Context) error {
	var productName dto.SearchName
	err := c.Bind(&productName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.StatusBadRequest("bad request.", err))
	}

	product, err := h.service.GetProductsByName(productName.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusInternalServerError("internal server error.", err))
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, response.StatusNotFound("product not found", nil))
	}

	return c.JSON(http.StatusOK, response.StatusOKWithData("success.", product))
}
