package healthcheck

import (
	"github.com/danisbagus/shopping-cart-api/core/port"
)

type (
	Service struct {
		repo port.HealthCheckRepo
	}
)

func NewService(repo port.HealthCheckRepo) port.HealthCheckService {
	return &Service{
		repo: repo,
	}
}

func (s Service) Ping() map[string]bool {

	err := s.repo.Ping()
	if err != nil {
		return map[string]bool{"ping": false}
	}

	return map[string]bool{"ping": true}
}
