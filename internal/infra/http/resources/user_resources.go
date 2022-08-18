package resources

import (
	"startUp/internal/domain"
	"time"
)

type UserDto struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Role        uint8     `json:"role"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	DeletedDate time.Time `json:"deleted_date"`
}

func MapDomainToUserDto(user *domain.User) *UserDto {
	return &UserDto{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Role:        uint8(user.Role),
		CreatedDate: user.CreatedDate,
		UpdatedDate: user.UpdatedDate,
		DeletedDate: user.DeletedDate,
	}
}

type TokenDto struct {
	Token string `json:"token"`
}

func MapDomainTokenDto(token string) *TokenDto {
	return &TokenDto{
		Token: token,
	}
}

func MapDomainToUserDtoCollection(users []domain.User) []UserDto {
	var result []UserDto
	for _, t := range users {
		dto := MapDomainToUserDto(&t)
		result = append(result, *dto)
	}
	return result
}
