package config

import (
	"github.com/danisbagus/shopping-cart-api/utils/config/logger"
	"github.com/danisbagus/shopping-cart-api/utils/config/psql"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	godotenv.Load()
	psql.Init()
	logger.Init()

}
