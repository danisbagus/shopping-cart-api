package api

import (
	healthCheckRepo "github.com/danisbagus/shopping-cart-api/infrastructure/repo/psql/healthCheck"
	userRepo "github.com/danisbagus/shopping-cart-api/infrastructure/repo/psql/user"

	healthCheckService "github.com/danisbagus/shopping-cart-api/core/healthCheck"
	userService "github.com/danisbagus/shopping-cart-api/core/user"

	healthCheckHandlerV1 "github.com/danisbagus/shopping-cart-api/interface/api/v1/healthCheck"
	userHandlerV1 "github.com/danisbagus/shopping-cart-api/interface/api/v1/user"

	"github.com/danisbagus/shopping-cart-api/utils/config/psql"

	"github.com/labstack/echo/v4"
)

func API(route *echo.Echo) {
	// db
	db := psql.GetDB()

	// repo
	healthCheckRepo := healthCheckRepo.NewRepo(db)
	userRepo := userRepo.NewRepo(db)

	// service
	healthCheckService := healthCheckService.NewService(healthCheckRepo)
	userService := userService.NewService(userRepo)

	// handler V1
	healthCheckHandlerV1 := healthCheckHandlerV1.NewHandler(healthCheckService)
	userHandlerV1 := userHandlerV1.NewHandler(userService)

	// route v1
	healthCheckRouterV1 := route.Group("/v1/health-check")
	healthCheckRouterV1.GET("/ping", healthCheckHandlerV1.Ping)

	userRouterV1 := route.Group("/v1/user")
	userRouterV1.POST("/register/customer", userHandlerV1.RegisterCustomer)
	userRouterV1.POST("/login", userHandlerV1.Login)
}
