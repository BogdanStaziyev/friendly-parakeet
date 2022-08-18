package domain

import "time"

type Role int

const (
	ROLE_ADMIN Role = 1
	// ROLE_USER ROLE_CLIENT
	ROLE_USER = 2
)

type User struct {
	Id          int64
	Name        string
	Email       string
	Password    string
	Passhash    []byte
	Role        Role
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate time.Time
}
