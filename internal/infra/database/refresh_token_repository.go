package database

import (
	"fmt"
	"time"

	"github.com/upper/db/v4"
	"startUp/internal/domain"
)

type refreshToken struct {
	Id          int64      `db:"id,omitempty"`
	UserId      int64      `db:"user_id"`
	Token       string     `db:"token"`
	ExpireDate  *time.Time `db:"expire_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type refreshTokenAndUser struct {
	refreshToken `db:",inline"`
	user         `db:", inline"`
}

type RefreshTokensRepository interface {
	Save(token *domain.RefreshToken) (domain.Role, *domain.RefreshToken, error)
	FindOne(id int64) (*domain.RefreshToken, error)
	Update(token *domain.RefreshToken) error
	Delete(userId, tokenId int64) error
}

type tokenRepository struct {
	coll    db.Collection
	session *db.Session
}

func NewRefreshTokenRepository(dbSession *db.Session) RefreshTokensRepository {
	return &tokenRepository{
		coll:    (*dbSession).Collection("refresh_tokens"),
		session: dbSession,
	}
}

func (r *tokenRepository) Save(token *domain.RefreshToken) (domain.Role, *domain.RefreshToken, error) {
	tkn := mapRefreshTokenDomainToDbModel(token)

	err := r.coll.InsertReturning(tkn)
	if err != nil {
		return 0, nil, fmt.Errorf("sessionTokensRepository SaveSessionToken: %w", err)
	}

	return token.UserRole, mapRefreshTokenDbModelToDomain(tkn), nil
}

func (r *tokenRepository) FindOne(tokenId int64) (*domain.RefreshToken, error) {

	req := r.coll.Session().SQL().Select(
		db.Raw("t.*"),
		db.Raw("u.*"),
	).From("refresh_tokens t").
		LeftJoin("users u").On("t.user_id = u.id").
		Where(
			db.Cond{"t.id": tokenId},
			db.Cond{"expire_date >": time.Now()},
			db.Cond{"t.deleted_date IS": nil},
			db.Cond{"u.deleted_date IS": nil},
		)

	data := refreshTokenAndUser{}
	err := req.One(&data)
	if err != nil {
		return nil, fmt.Errorf("sessionTokensRepository FindOneSessionToken: %w", err)
	}
	data.refreshToken.Id = tokenId
	return mapRefreshTokenAndUserDbModelToDomain(&data), nil
}

func (r *tokenRepository) Update(token *domain.RefreshToken) error {
	tkn := mapRefreshTokenDomainToDbModel(token)

	err := r.coll.Find(tkn.Id).Update(tkn)
	if err != nil {
		return fmt.Errorf("sessionTokensRepository UpdateSessionToken: %w", err)
	}

	return nil
}

func (r *tokenRepository) Delete(userId, tokenId int64) error {
	userIdCond := db.Cond{"user_id": userId}
	tokenIdCond := db.Cond{"id": tokenId}
	err := r.coll.Find(db.And(userIdCond, tokenIdCond)).Update(map[string]interface{}{"deleted_date": time.Now()})
	if err != nil {
		return fmt.Errorf("sessionTokensRepository DeleteSessionToken: %w", err)
	}

	return nil
}

func mapRefreshTokenDbModelToDomain(tkn *refreshToken) *domain.RefreshToken {
	return &domain.RefreshToken{
		Id:          tkn.Id,
		UserId:      tkn.UserId,
		Token:       tkn.Token,
		ExpireDate:  getTimeFromTimePtr(tkn.ExpireDate),
		DeletedDate: getTimeFromTimePtr(tkn.DeletedDate),
	}
}

func mapRefreshTokenAndUserDbModelToDomain(model *refreshTokenAndUser) *domain.RefreshToken {
	refreshToken := mapRefreshTokenDbModelToDomain(&model.refreshToken)
	refreshToken.User = mapUserDbModelToDomain(&model.user)
	return refreshToken
}

func mapRefreshTokenDomainToDbModel(tkn *domain.RefreshToken) *refreshToken {
	return &refreshToken{
		Id:          tkn.Id,
		UserId:      tkn.UserId,
		Token:       tkn.Token,
		ExpireDate:  getTimePtrFromTime(tkn.ExpireDate),
		DeletedDate: getTimePtrFromTime(tkn.DeletedDate),
	}
}
