package domain

import (
	"time"
)

type (
	User struct {
		ID        uint64
		Name      string
		Email     string
		Password  string
		RoleID    uint64
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	UserAuth struct {
		User
		AccessToken string
	}
)
