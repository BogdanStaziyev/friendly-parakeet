package domain

import "time"

type RefreshToken struct {
	Id          int64
	UserId      int64
	UserRole    Role
	Token       string
	ExpireDate  time.Time
	DeletedDate time.Time
	*User
}
