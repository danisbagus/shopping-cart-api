package user

import (
	"errors"
	"fmt"
	"testing"

	"github.com/danisbagus/shopping-cart-api/core/domain"
	userRepo "github.com/danisbagus/shopping-cart-api/infrastructure/repo/mock/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockRepo = &userRepo.Repo{Mock: mock.Mock{}}
var service = NewService(mockRepo)

func TestRegisterCustomer(t *testing.T) {

	// TABLE TEST

	type param struct {
		form *domain.User
	}

	type expected struct {
		userAuth *domain.UserAuth
		err      error
	}

	tests := []struct {
		testName string
		param    param
		expected expected
	}{
		{
			testName: "Register error FindOneByEmail",
			param: param{form: &domain.User{
				Email: "error@testing.com",
			}},
			expected: expected{
				userAuth: nil,
				err:      errors.New("FindOneByEmail error"),
			},
		},
		{
			testName: "Register email already used",
			param: param{form: &domain.User{
				Email: "used@testing.com",
			}},
			expected: expected{
				userAuth: nil,
				err:      errors.New("email already used"),
			},
		},
		{
			testName: "Register error Insert",
			param: param{form: &domain.User{
				Password: "user1",
				Email:    "user1@testing.com",
			}},
			expected: expected{
				userAuth: nil,
				err:      errors.New("Insert error"),
			},
		},
		{
			testName: "Register success",
			param: param{form: &domain.User{
				Password: "user2",
				Email:    "user2@testing.com",
			}},
			expected: expected{
				userAuth: &domain.UserAuth{
					User: domain.User{
						Password: "user2",
						Email:    "user2@testing.com",
					},
				},
				err: nil,
			},
		},
	}

	mockRepo.Mock.On("FindOneByEmail", "error@testing.com").Return(nil, errors.New("FindOneByEmail error"))

	mockRepo.Mock.On("FindOneByEmail", "used@testing.com").Return(&domain.User{ID: 1}, nil)

	mockRepo.Mock.On("FindOneByEmail", "user1@testing.com").Return(&domain.User{}, nil)

	mockRepo.Mock.On("FindOneByEmail", "user2@testing.com").Return(&domain.User{}, nil)

	mockRepo.Mock.On("Insert", mock.MatchedBy(func(user *domain.User) bool { return user.Email == "user1@testing.com" })).Return(nil, errors.New("Insert error"))

	mockRepo.Mock.On("Insert", mock.MatchedBy(func(user *domain.User) bool { return user.Email == "user2@testing.com" })).Return(&domain.User{
		Password: "user2",
		Email:    "user2@testing.com",
	}, nil)

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			userAuth, err := service.Register(test.param.form)

			assert.Equal(t, test.expected.err, err)

			if userAuth != nil {
				assert.Equal(t, test.expected.userAuth.Email, userAuth.Email)
				assert.Equal(t, test.expected.userAuth.Password, userAuth.Password)
			}
		})
	}
}

func TestLogin(t *testing.T) {

	// SUB TEST

	mockRepo.Mock.On("FindOneByEmail", "error@testing.com").Return(nil, errors.New("FindOneByEmail error"))

	mockRepo.Mock.On("FindOneByEmail", "user1@testing.com").Return(&domain.User{}, nil)

	mockRepo.Mock.On("FindOneByEmail", "user3@testing.com").Return(&domain.User{ID: 3, Password: "$2a$14$vftRog9cw/awljTJggIkGeYvMWi3oKt3hqUFVT2k84L/AHvbHZUbK"}, nil)

	t.Run("Insert error FindOneByEmail", func(t *testing.T) {
		userAuth, err := service.Login("error@testing.com", "error")
		assert.Nil(t, userAuth)
		assert.Equal(t, errors.New("FindOneByEmail error"), err)
	})

	t.Run("Insert error user not found", func(t *testing.T) {
		userAuth, err := service.Login("user1@testing.com", "user1")
		assert.Nil(t, userAuth)
		assert.Equal(t, errors.New("user not found"), err)
	})

	t.Run("Insert error user not found", func(t *testing.T) {
		userAuth, err := service.Login("user3@testing.com", "user1")
		assert.Nil(t, userAuth)
		assert.Equal(t, errors.New("invalid credential"), err)
	})

	t.Run("Insert success", func(t *testing.T) {
		userAuth, err := service.Login("user3@testing.com", "user3")
		assert.NotNil(t, userAuth)
		assert.Equal(t, nil, err)
	})
}

func TestHashPassword(t *testing.T) {
	hasPassword, _ := domain.HashPassword("user3")
	fmt.Println(hasPassword)
}
