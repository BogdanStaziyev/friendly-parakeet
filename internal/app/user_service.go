package app

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"startUp/internal/domain"
	"startUp/internal/infra/database"
)

type UserService interface {
	Save(user *domain.User) (*domain.User, error)
	FindOne(userId int64, q *domain.UrlQueryParams) (*domain.User, error)
	FindAll(q *domain.UrlQueryParams) ([]domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(userId int64) error
	LogIn(user *domain.User) (*domain.User, error)
}

type userService struct {
	userRepo *database.UserRepository
}

func NewUserService(u *database.UserRepository) UserService {
	return &userService{
		userRepo: u,
	}
}

func (u *userService) Save(user *domain.User) (*domain.User, error) {
	// get hash password
	phash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 31121991)
	if err != nil {
		return nil, fmt.Errorf("userService SaveUser: %w", err)
	}
	user.Passhash = phash
	// save user
	storedUser, err := (*u.userRepo).Save(user)
	if err != nil {
		return nil, fmt.Errorf("userService SaveUser: %w", err)
	}

	return storedUser, nil
}

func (u *userService) FindOne(userId int64, q *domain.UrlQueryParams) (*domain.User, error) {
	storedUser, err := (*u.userRepo).FindOne(userId, q)
	if err != nil {
		return nil, fmt.Errorf("userService FindOne: %w", err)
	}
	return storedUser, err
}

func (u *userService) FindAll(q *domain.UrlQueryParams) ([]domain.User, error) {
	users, err := (*u.userRepo).FindAll(q)
	if err != nil {
		return nil, fmt.Errorf("userService FindAll: %w", err)
	}

	return users, err
}

func (u *userService) Update(user *domain.User) (*domain.User, error) {
	updateUser, err := (*u.userRepo).Update(user)
	if err != nil {
		return nil, fmt.Errorf("userService UpdateUser: %w", err)
	}

	return updateUser, nil
}

func (u *userService) Delete(userId int64) error {
	err := (*u.userRepo).Delete(userId)
	if err != nil {
		return fmt.Errorf("userService DeleteUser: %w", err)
	}
	return nil
}

func (u *userService) LogIn(userQueried *domain.User) (*domain.User, error) {
	userStored, err := (*u.userRepo).FindOneByEmail(userQueried.Email, nil)
	if err != nil {
		return nil, fmt.Errorf("userService LogInUser: %w", err)
	}
	err = CheckPassword(userQueried, userStored)
	if err != nil {
		return nil, fmt.Errorf("userService LogInUser: %w", err)
	}
	return userStored, nil
}

// CheckPassword return nil on success or an error on failure
func CheckPassword(userQueried, userStored *domain.User) error {
	return bcrypt.CompareHashAndPassword(userStored.Passhash, []byte(userQueried.Password))
}
