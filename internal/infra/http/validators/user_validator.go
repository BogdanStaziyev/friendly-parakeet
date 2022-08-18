package validators

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"net/http"
	"startUp/internal/domain"
)

type UserValidator struct {
	validator *validator.Validate
}

func NewUserValidator() *UserValidator {
	return &UserValidator{
		validator: validator.New(),
	}
}

func (u UserValidator) ValidateAndMap(r *http.Request) (*domain.User, error) {
	var userResource UserRequest
	err := json.NewDecoder(r.Body).Decode(&userResource)
	if err != nil {
		return nil, fmt.Errorf("userValidator ValideteAndMap %w", err)
	}

	err = u.validator.Struct(userResource)
	if err != nil {
		return nil, fmt.Errorf("userValidator ValideteAndMap %w", err)
	}
	return mapUserRequestToDomain(&userResource), nil
}

type UserLoginValidator struct {
	validator *validator.Validate
}

func NewUserLoginValidator() *UserLoginValidator {
	return &UserLoginValidator{
		validator: validator.New(),
	}
}

func (u UserLoginValidator) ValidatorAndMap(r *http.Request) (*domain.User, error) {
	var userResource userLoginRequest
	err := json.NewDecoder(r.Body).Decode(&userResource)
	if err != nil {
		return nil, fmt.Errorf("LoginValidator VAlidateAndMap: %w", err)
	}
	err = u.validator.Struct(userResource)
	if err != nil {
		return nil, fmt.Errorf("userLoginValidator VAlidateAndMap: %w", err)
	}
	return mapUserLoginRequestToDomain(&userResource), nil
}

type UserUpdateValidator struct {
	validator *validator.Validate
}

func NewUserUpdateValidator() *UserUpdateValidator {
	return &UserUpdateValidator{
		validator: validator.New(),
	}
}

func (u UserUpdateValidator) ValidateAndMap(r *http.Request) (*domain.User, error) {
	var userResource userUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&userResource)
	if err != nil {
		return nil, fmt.Errorf("useerUpdateValidator VAlidateAndMap: %w", err)
	}
	err = u.validator.Struct(userResource)
	if err != nil {
		return nil, fmt.Errorf("useerUpdateValidator VAlidateAndMap: %w", err)
	}
	return mapUserUpdateRequestDomain(&userResource), nil
}
