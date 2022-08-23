package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"startUp/internal/app"
	"startUp/internal/domain"
)

type contextUserKey string

const contextUserIdKey contextUserKey = "1"
const BEARER_SCHEMA = "Bearer "

func AuthMiddleware(s app.RefreshTokenService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//verify access token
			if user, err := authorizeWithAccessToken(r, s); err == nil {
				// add user object to the request context
				r = r.WithContext(context.WithValue(r.Context(), contextUserIdKey, user))
				// serve next
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
		})
	}
}

func AdminOnli(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetAuthorizedUser(r)
		if user == nil {
			log.Println("Warning! User authorization check is turned off!")
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	})
}

func authorizeWithAccessToken(r *http.Request, s app.RefreshTokenService) (*domain.RefreshToken, error) {
	//get jwt access token
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("failed to get token from Autorization header")
	}
	token := authHeader[len(BEARER_SCHEMA):]

	//verify jwt token
	user, err := s.VerifyAccessToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to get token from Autorization header")
	}
	return user, nil
}

//return authorized user object or nil if user is not authorized
func GetAuthorizedUser(r *http.Request) *domain.RefreshToken {
	object := r.Context().Value(contextUserIdKey)
	if user, ok := object.(*domain.RefreshToken); ok {
		return user
	}
	return nil
}
