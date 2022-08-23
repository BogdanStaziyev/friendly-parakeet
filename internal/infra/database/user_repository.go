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
	UpdatedDate time.Time  `db:"updated_date, omitempty"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	FindOne(id int64, q *domain.UrlQueryParams) (*domain.User, error)
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

func (r *userRepository) Save(user *domain.User) (*domain.User, error) {
	modelUser := mapDomainToUserDbModel(user)

	err := r.coll.InsertReturning(modelUser)
	if err != nil {
		return nil, fmt.Errorf("userRepository SaveUser: %w", err)
	}

	return mapUserDbModelToDomain(modelUser), nil
}

func (r *userRepository) FindOne(userId int64, q *domain.UrlQueryParams) (*domain.User, error) {
	var storedUser *user
	query := mapDomainToDbQueryParams(q)
	err := query.ApplyToResult(r.coll.Find(userId)).One(&storedUser)
	if err != nil {
		return nil, fmt.Errorf("userRepository FindOneUser: %w", err)
	}

	return mapUserDbModelToDomain(storedUser), nil
}

func (r *userRepository) FindOneByEmail(userEmail string, q *domain.UrlQueryParams) (*domain.User, error) {
	var storedUser *user
	query := mapDomainToDbQueryParams(q)
	err := query.ApplyToResult(r.coll.Find(db.Cond{"users.email": userEmail})).One(&storedUser)
	if err != nil {
		return nil, fmt.Errorf("userRepository FindOneUserWithRole: %w", err)
	}

	return mapUserDbModelToDomain(storedUser), nil
}

func (r *userRepository) FindAll(q *domain.UrlQueryParams) ([]domain.User, error) {
	var storedUsers []user
	query := mapDomainToDbQueryParams(q)
	err := query.ApplyToResult(r.coll.Find()).OrderBy("-created_date").All(&storedUsers)
	if err != nil {
		return nil, fmt.Errorf("userRepository PaginateAllUsers: %w", err)
	}

	return mapUserDbModelToDomainCollection(storedUsers), nil
}

func (r *userRepository) Update(user *domain.User) (*domain.User, error) {
	userToUpdate := mapDomainToUserDbModel(user)
	userToUpdate.UpdatedDate = time.Now()
	err := r.coll.UpdateReturning(userToUpdate)
	if err != nil {
		return nil, fmt.Errorf("userRepository UpdateUser: %w", err)
	}

	return mapUserDbModelToDomain(userToUpdate), nil
}

func (r *userRepository) Delete(userId int64) error {
	err := r.coll.Find(userId).Update(map[string]interface{}{"deleted_date": time.Now()})
	if err != nil {
		return fmt.Errorf("userRepository DeleteUser: %w", err)
	}

	return nil
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

func mapUserDbModelToDomainCollection(users []user) []domain.User {
	var result []domain.User

	for _, t := range users {
		newUsers := mapUserDbModelToDomain(&t)
		result = append(result, *newUsers)
	}
	return result
}

func mapDomainToUserDbModel(usr *domain.User) *user {
	return &user{
		Id:          usr.Id,
		Name:        usr.Name,
		Email:       usr.Email,
		Role:        uint8(usr.Role),
		Passhash:    usr.Passhash,
		CreatedDate: usr.CreatedDate,
		UpdatedDate: usr.UpdatedDate,
		DeletedDate: getTimePtrFromTime(usr.DeletedDate),
	}
}
