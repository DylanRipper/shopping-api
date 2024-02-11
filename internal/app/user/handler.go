package user

import (
	"net/http"
	"shopping-api/internal/dto"
	"shopping-api/internal/factory"
	middlewares "shopping-api/internal/middleware"
	"shopping-api/internal/model"
	"shopping-api/internal/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) CreateUser(c echo.Context) error {
	user := model.Users{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.StatusBadRequest("bad request.", err))

	}

	err = h.service.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.StatusBadRequest("bad request.", err))

	}

	return c.JSON(http.StatusOK, response.StatusOK("successfully created user."))
}

func (h *handler) GetUser(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	user, err := h.service.GetUser(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusInternalServerError, response.StatusInternalServerError("internal server error", err))

	} else if err == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, response.StatusNotFound("not found.", "data user not found"))
	}

	return c.JSON(http.StatusOK, response.StatusOKWithData("success.", user))
}

func (h *handler) DeleteUser(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	err := h.service.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusInternalServerError("internal server error", err))

	}

	return c.JSON(http.StatusOK, response.StatusOK("success delete user."))
}

func (h *handler) LoginUser(c echo.Context) error {
	user := dto.UserLogin{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.StatusBadRequest("bad request.", err))
	}

	data, err := h.service.LoginUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusInternalServerError("internal server error", err))
	}

	return c.JSON(http.StatusOK, response.StatusOKWithData("success.", data))
}
