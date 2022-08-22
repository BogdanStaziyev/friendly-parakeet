package database

import (
	"fmt"
	"github.com/upper/db/v4"
	"startUp/internal/domain"
	"time"
)

type refreshToken struct {
	Id          int64      `json:"id,omitempty"`
	UserId      int64      `json:"user_id"`
	Token       string     `json:"token"`
	ExpireDate  *time.Time `json:"expire_date"`
	DeletedDate *time.Time `json:"deleted_date"`
}

type refreshTokenAndUser struct {
	refreshToken `db:",inline"`
	user         `db:", inline"`
}

type RefreshTokenRepository interface {
	Save(token *domain.RefreshToken) (domain.Role, *domain.RefreshToken, error)
	FindOne(tokenId int64) (*domain.RefreshToken, error)
	Update(token *domain.RefreshToken) error
	Delete(userId, tokenId int64) error
}

type tokenRepository struct {
	coll    db.Collection
	session *db.Session
}

func NewRefreshTokenRepository(dbSession *db.Session) RefreshTokenRepository {
	return &tokenRepository{
		coll:    (*dbSession).Collection("refresh_tokens"),
		session: dbSession,
	}
}

func (t *tokenRepository) Save(token *domain.RefreshToken) (domain.Role, *domain.RefreshToken, error) {
	tkn := mapRefreshTokenDomainToDbModel(token)

	err := t.coll.InsertReturning(tkn)
	if err != nil {
		return 0, nil, fmt.Errorf("sessionTokenRepository SaveSessionToken: %w", err)
	}

	return token.UserRole, mapRefreshTokenDbModelToDomain(tkn), nil
}

func (t *tokenRepository) FindOne(tokenId int64) (*domain.RefreshToken, error) {
	req := t.coll.Session().SQL().Select(
		db.Raw("t.*"),
		db.Raw("u.*"),
	).From("refresh_token t").
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
		return nil, fmt.Errorf("sessionTokenRepository FindOneSessionToken: %w", err)
	}
	data.refreshToken.Id = tokenId
	return mapRefreshTokenAndUserDbModelToDomain(&data), nil
}

func (t *tokenRepository) Update(token *domain.RefreshToken) error {
	tkn := mapRefreshTokenDomainToDbModel(token)

	err := t.coll.Find(tkn.Id).Update(tkn)
	if err != nil {
		return fmt.Errorf("sessionTokensRepository UpdateSessionToken: %w", err)
	}

	return nil
}

func (t *tokenRepository) Delete(userId, tokenId int64) error {
	userIdCond := db.Cond{"user_id": userId}
	tokenIdCond := db.Cond{"id": tokenId}
	err := t.coll.Find(db.And(userIdCond, tokenIdCond)).Update(map[string]interface{}{"deleted_date": time.Now()})
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
