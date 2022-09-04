package user

import (
	"net/http"

	"github.com/danisbagus/shopping-cart-api/core/domain"
	"github.com/danisbagus/shopping-cart-api/core/port"
	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
		service port.UserService
	}
)

func NewHandler(service port.UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) RegisterCustomer(c echo.Context) error {

	req := new(RequestRegisterCustomer)

	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = req.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	form := new(domain.User)
	form.Name = req.Name
	form.Email = req.Email
	form.Password = req.Password

	user, err := h.service.Register(form)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := NewResponseUserAuth("successfully register customer", user)
	return c.JSON(http.StatusOK, res)
}

func (h Handler) Login(c echo.Context) error {

	req := new(RequestLogin)

	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = req.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := NewResponseUserAuth("successfully login", user)
	return c.JSON(http.StatusOK, res)
}
