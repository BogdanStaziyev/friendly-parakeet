package validators

import (
	"startUp/internal/domain"
)

type UserRequest struct {
	Name     string `json:"name" validate:"required,gte=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	Role     uint8  `json:"role" validate:"required,numeric"`
}

func mapUserRequestToDomain(request *UserRequest) *domain.User {
	return &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     domain.Role(request.Role),
	}
}

type userLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func mapUserLoginRequestToDomain(request *userLoginRequest) *domain.User {
	return &domain.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

type userUpdateRequest struct {
	Name  string `json:"name" validate:"required,gte=3"`
	Email string `json:"email" validate:"required,email"`
	Role  uint8  `json:"role" validate:"required,numeric"`
}

func mapUserUpdateRequestDomain(request *userUpdateRequest) *domain.User {
	return &domain.User{
		Name:  request.Name,
		Email: request.Email,
		Role:  domain.Role(request.Role),
	}
}
