package database

import (
	"fmt"
	"time"

	"github.com/upper/db/v4"
	"startUp/internal/domain"
)

const UserTable = "users"

type user struct {
	Id          int64      `db:"id,omitempty"`
	Name        string     `db:"name"`
	Email       string     `db:"email"`
	Passhash    []byte     `db:"passhash,omitempty"`
	Role        uint8      `db:"role_id"`
	CreatedDate time.Time  `db:"created_date,omitempty"`
	UpdatedDate time.Time  `db:"updated_date,omitempty"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	FindOne(userId int64, q *domain.UrlQueryParams) (*domain.User, error)
	FindOneByEmail(email string, q *domain.UrlQueryParams) (*domain.User, error)
	FindAll(q *domain.UrlQueryParams) ([]domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(userId int64) error
}

type userRepository struct {
	coll    db.Collection
	session *db.Session
}

func NewUserRepository(dbSession *db.Session) UserRepository {
	return &userRepository{
		coll:    (*dbSession).Collection(UserTable),
		session: dbSession,
	}
}

func (u *userRepository) Save(user *domain.User) (*domain.User, error) {
	modelUser := mapDomainToUserModel(user)

	err := u.coll.InsertReturning(modelUser)
	if err != nil {
		return nil, fmt.Errorf("userRepository SaveUser: %w", err)
	}

	return mapUserDbModelToDomain(modelUser), nil
}

func (u *userRepository) FindOne(userId int64, q *domain.UrlQueryParams) (*domain.User, error) {
	var modelUser *user
	query := mapDomainToDbQueryParams(q)
	err := query.ApplyToResult(u.coll.Find(userId)).One(&modelUser)
	if err != nil {
		return nil, fmt.Errorf("userRepository FindOne: %w", err)
	}

	return mapUserDbModelToDomain(modelUser), nil
}

func (u *userRepository) FindOneByEmail(email string, q *domain.UrlQueryParams) (*domain.User, error) {
	var modelUser *user
	query := mapDomainToDbQueryParams(q)
	err := query.ApplyToResult(u.coll.Find(db.Cond{"users.email": email})).One(&modelUser)
	if err != nil {
		return nil, fmt.Errorf("userRepository FindOneUserByEmail: %w", err)
	}

	return mapUserDbModelToDomain(modelUser), nil
}

func (u *userRepository) FindAll(q *domain.UrlQueryParams) ([]domain.User, error) {
	var modelUsers []user
	query := mapDomainToDbQueryParams(q)
	err := query.ApplyToResult(u.coll.Find()).OrderBy("-created_date").All(&modelUsers)
	if err != nil {
		return nil, fmt.Errorf("userRepository FindAll: %w", err)
	}

	return mapUserIDmModelToDomainCollection(modelUsers), nil
}

func (u *userRepository) Update(user *domain.User) (*domain.User, error) {
	userToUpdate := mapDomainToUserModel(user)
	userToUpdate.UpdatedDate = time.Now()
	err := u.coll.UpdateReturning(userToUpdate)
	if err != nil {
		return nil, fmt.Errorf("userRepository Update: %w", err)
	}

	return mapUserDbModelToDomain(userToUpdate), nil
}

func (u *userRepository) Delete(id int64) error {
	err := u.coll.Find(id).Update(map[string]interface{}{"deleted_date": time.Now()})
	if err != nil {
		return fmt.Errorf("userRepository DeleteUser: %w", err)
	}

	return nil
}

func mapDomainToUserModel(usr *domain.User) *user {
	return &user{
		Id:          usr.Id,
		Name:        usr.Name,
		Email:       usr.Email,
		Passhash:    usr.Passhash,
		Role:        uint8(usr.Role),
		CreatedDate: usr.CreatedDate,
		UpdatedDate: usr.UpdatedDate,
		DeletedDate: getTimePtrFromTime(usr.DeletedDate),
	}
}

func mapUserDbModelToDomain(usr *user) *domain.User {
	return &domain.User{
		Id:          usr.Id,
		Name:        usr.Name,
		Email:       usr.Email,
		Passhash:    usr.Passhash,
		Role:        domain.Role(usr.Role),
		CreatedDate: usr.CreatedDate,
		UpdatedDate: usr.UpdatedDate,
		DeletedDate: getTimeFromTimePtr(usr.DeletedDate),
	}
}

func mapUserIDmModelToDomainCollection(users []user) []domain.User {
	var result []domain.User

	for _, t := range users {
		newUsers := mapUserDbModelToDomain(&t)
		result = append(result, *newUsers)
	}

	return result
}
