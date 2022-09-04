package user

import (
	"fmt"
	"time"

	"github.com/danisbagus/shopping-cart-api/core/domain"
	"github.com/danisbagus/shopping-cart-api/core/port"
	"github.com/danisbagus/shopping-cart-api/utils/constant"
)

type (
	Service struct {
		repo port.UserRepo
	}
)

func NewService(repo port.UserRepo) port.UserService {
	return &Service{
		repo: repo,
	}
}

func (s Service) Register(form *domain.User) (*domain.UserAuth, error) {

	user, err := s.repo.FindOneByEmail(form.Email)
	if err != nil {
		return nil, err
	}

	if user.ID != 0 {
		return nil, fmt.Errorf("email already used")
	}

	// hashing password
	hashPassword, _ := domain.HashPassword(form.Password)

	form.Password = hashPassword
	form.RoleID = constant.RoleIDCustomer
	form.CreatedAt = time.Now()
	form.UpdatedAt = time.Now()

	insertData, err := s.repo.Insert(form)
	if err != nil {
		return nil, err
	}

	// generate access token
	accessToken, err := domain.GenerateAccessToken(insertData.ID, insertData.RoleID)
	if err != nil {
		return nil, err
	}

	userAuth := new(domain.UserAuth)
	userAuth.ID = insertData.ID
	userAuth.Name = insertData.Name
	userAuth.Email = insertData.Email
	userAuth.Password = insertData.Password
	userAuth.RoleID = insertData.RoleID
	userAuth.AccessToken = accessToken

	return userAuth, nil
}

func (s Service) Login(email, password string) (*domain.UserAuth, error) {

	user, err := s.repo.FindOneByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	// check password
	isValidPassword := domain.CheckPasswordHash(password, user.Password)
	if !isValidPassword {
		return nil, fmt.Errorf("invalid credential")
	}

	// generate access token
	accessToken, err := domain.GenerateAccessToken(user.ID, user.RoleID)
	if err != nil {
		return nil, err
	}

	userAuth := new(domain.UserAuth)
	userAuth.ID = user.ID
	userAuth.Name = user.Name
	userAuth.Email = user.Email
	userAuth.Password = user.Password
	userAuth.RoleID = user.RoleID
	userAuth.AccessToken = accessToken

	return userAuth, nil
}
