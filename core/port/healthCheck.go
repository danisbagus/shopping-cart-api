package port

type (
	HealthCheckRepo interface {
		Ping() error
	}

	HealthCheckService interface {
		Ping() map[string]bool
	}
)
