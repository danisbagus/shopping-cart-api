// Code generated by mockery v2.10.0. DO NOT EDIT.

package user

import (
	domain "github.com/danisbagus/shopping-cart-api/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repo is an autogenerated mock type for the UserRepo type
type Repo struct {
	mock.Mock
}

// FindOneByEmail provides a mock function with given fields: email
func (_m *Repo) FindOneByEmail(email string) (*domain.User, error) {
	ret := _m.Called(email)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0
func (_m *Repo) Insert(_a0 *domain.User) (*domain.User, error) {
	ret := _m.Called(_a0)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(*domain.User) *domain.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}