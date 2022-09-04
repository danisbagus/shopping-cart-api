package user

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	RequestRegisterCustomer struct {
		Name            string `json:"name"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	RequestLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (r RequestRegisterCustomer) Validate() error {
	if err := validation.Validate(r.Name, validation.Required); err != nil {
		return errors.New("name is required")
	} else if err := validation.Validate(r.Email, validation.Required); err != nil {
		return errors.New("email is required")
	} else if err := validation.Validate(r.Password, validation.Required); err != nil {
		return errors.New("password is required")
	} else if r.Password != r.ConfirmPassword {
		return errors.New("invalid confirm password")
	}

	return nil
}

func (r RequestLogin) Validate() error {
	if err := validation.Validate(r.Email, validation.Required); err != nil {
		return errors.New("email is required")
	} else if err := validation.Validate(r.Password, validation.Required); err != nil {
		return errors.New("password is required")
	}

	return nil
}
