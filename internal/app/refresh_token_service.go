package app

import (
	"crypto/rand"
	"fmt"
	"startUp/internal/domain"
	"startUp/internal/infra/database"
	"time"

	"github.com/golang-jwt/jwt"
)

type RefreshTokenService interface {
	CreateAccessToken(storedToken *domain.RefreshToken) (string, error)
	VerifyAccessToken(tokenString string) (*domain.RefreshToken, error)
	DeleteSessionToken(userId, tokenId int64) error
	CreateRefreshToken(user *domain.User) (*domain.RefreshToken, error)
}

type sessionService struct {
	sessionRepo  *database.RefreshTokenRepository
	secretAccess []byte
}

type userAccessClaims struct {
	UserId   int64       `json:"user_id"`
	UserRole domain.Role `json:"user_role"`
	TokenId  int64       `json:"token_id"`
	jwt.StandardClaims
}

func NewRefreshTokenService(u *database.RefreshTokenRepository, secretAccess []byte) RefreshTokenService {
	return &sessionService{
		sessionRepo:  u,
		secretAccess: secretAccess,
	}
}

func (s *sessionService) CreateAccessToken(storedToken *domain.RefreshToken) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, userAccessClaims{
		UserId:   storedToken.UserId,
		UserRole: storedToken.UserRole,
		TokenId:  storedToken.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: getNewAccessExpireUnixTime(),
		},
	})

	return jwtToken.SignedString(s.secretAccess)
}

func (s *sessionService) VerifyAccessToken(tokenString string) (*domain.RefreshToken, error) {

	claims, err := parseJwt(tokenString, &userAccessClaims{}, s.secretAccess)
	if err != nil {
		return nil, fmt.Errorf("sessionService VerifyAccessToken: %w", err)
	}

	accessClaims := claims.(*userAccessClaims)
	return &domain.RefreshToken{
		Id:       accessClaims.TokenId,
		UserId:   accessClaims.UserId,
		UserRole: accessClaims.UserRole,
	}, nil
}

func (s *sessionService) DeleteSessionToken(userId, tokenId int64) error {
	return (*s.sessionRepo).Delete(userId, tokenId)
}

func (s *sessionService) CreateRefreshToken(user *domain.User) (*domain.RefreshToken, error) {
	newToken := domain.RefreshToken{
		UserId:     user.Id,
		UserRole:   user.Role,
		Token:      getNewRefreshToken(),
		ExpireDate: getNewRefreshExpireDate(),
	}

	role, storedToken, err := (*s.sessionRepo).Save(&newToken)
	if err != nil {
		return nil, fmt.Errorf("sessionService CreateRefreshToken: failed to save the token: %w", err)
	}
	storedToken.UserRole = role
	return storedToken, nil
}

func getNewRefreshToken() string {
	b := make([]byte, 20)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}

func getNewRefreshExpireDate() time.Time {
	return time.Now().AddDate(0, 0, 5) // expire after 5 days
}

func getNewAccessExpireUnixTime() int64 {
	return time.Now().Add(time.Hour * 4).Unix() //expire after 4 hour
}

func parseJwt(tokenString string, claims jwt.Claims, secret []byte) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("SessionService parseJWT: parse error: %w", err)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, fmt.Errorf("SessionService parseJWT: validation error: %w", err)
	}

	return claims, nil
}
