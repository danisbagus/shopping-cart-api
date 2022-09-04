package middleware

import (
	"strings"

	"github.com/danisbagus/shopping-cart-api/utils/config/logger"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

func APILogHandler(c echo.Context, req, res []byte) {
	logger.Info("incoming request",
		zap.String("request", string(req)),
		zap.String("response", string(res)),
		zap.Any("context", c),
	)
}

func APILogSkipper(c echo.Context) bool {
	paths := []string{}

	requestURI := c.Request().RequestURI
	for i := range paths {
		if strings.Contains(requestURI, paths[i]) {
			return true
		}
	}

	indexAlwaysAllowed := requestURI == "/" || requestURI == ""
	return indexAlwaysAllowed
}
