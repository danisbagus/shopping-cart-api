package healthcheck

import (
	"net/http"

	"github.com/danisbagus/shopping-cart-api/core/port"
	"github.com/danisbagus/shopping-cart-api/interface/api/common"
	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
		service port.HealthCheckService
	}
)

func NewHandler(service port.HealthCheckService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) Ping(c echo.Context) error {
	ping := h.service.Ping()

	response := new(common.DefaultResponse)
	response.SetResponseData("Success Ping", ping)
	return c.JSON(http.StatusOK, response)
}
