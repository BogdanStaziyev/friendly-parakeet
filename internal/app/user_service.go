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

func (s *userService) Save(user *domain.User) (*domain.User, error) {
	// get password hash
	phash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 2)
	if err != nil {
		return nil, fmt.Errorf("userService SaveUser: %w", err)
	}
	user.Passhash = phash
	// save user
	storedUser, err := (*s.userRepo).Save(user)
	if err != nil {
		return nil, fmt.Errorf("userService SaveUser: %w", err)
	}

	return storedUser, nil
}

func (s *userService) FindOne(userId int64, q *domain.UrlQueryParams) (*domain.User, error) {
	storedUser, err := (*s.userRepo).FindOne(userId, q)
	if err != nil {
		return nil, fmt.Errorf("userService FindOneUser: %w", err)
	}

	return storedUser, nil
}

func (s *userService) FindAll(q *domain.UrlQueryParams) ([]domain.User, error) {
	users, err := (*s.userRepo).FindAll(q)
	if err != nil {
		return nil, fmt.Errorf("userService FindAllUsers: %w", err)
	}

	return users, nil
}

func (s *userService) Update(user *domain.User) (*domain.User, error) {
	updatedUser, err := (*s.userRepo).Update(user)
	if err != nil {
		return nil, fmt.Errorf("userService UpdateUser: %w", err)
	}

	return updatedUser, nil
}

func (s *userService) Delete(userId int64) (err error) {
	err = (*s.userRepo).Delete(userId)
	if err != nil {
		return fmt.Errorf("userService DeleteUser: %w", err)
	}

	return nil
}

func (s *userService) LogIn(userQueried *domain.User) (*domain.User, error) {
	userStored, err := (*s.userRepo).FindOneByEmail(userQueried.Email, nil)
	if err != nil {
		return nil, fmt.Errorf("userService LogInUser: %w", err)
	}

	err = CheckPassword(userQueried, userStored)
	if err != nil {
		return nil, fmt.Errorf("userService LogInUser: %w", err)
	}

	return userStored, nil
}

// CheckPassword Return nil on success or an error on failure
func CheckPassword(userQueried, userStored *domain.User) error {
	return bcrypt.CompareHashAndPassword(userStored.Passhash, []byte(userQueried.Password))
}
