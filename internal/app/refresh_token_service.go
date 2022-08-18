package app

import "startUp/internal/domain"

type RefreshTokenService interface {
	CreateAccessToken(storedToken *domain.RefreshToken) (string, error)
	VerifyAccessToken(token string) (*domain.RefreshToken, error)
	DeleteSessionToken(userId, tokenId int64) error
	CreateRefreshToken(user *domain.User) (*domain.RefreshToken, error)
}
