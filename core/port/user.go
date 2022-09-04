package port

import (
	"github.com/danisbagus/shopping-cart-api/core/domain"
)

type (
	UserRepo interface {
		FindOneByEmail(email string) (*domain.User, error)
		Insert(user *domain.User) (*domain.User, error)
	}

	UserService interface {
		Register(user *domain.User) (*domain.UserAuth, error)
		Login(email, password string) (*domain.UserAuth, error)
	}
)
