package controllers

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"startUp/internal/app"
	"startUp/internal/infra/http/resources"
	"startUp/internal/infra/http/validators"
	"strconv"
)

const BEARER_SCHEMA = "Bearer "

type UserController struct {
	userService         *app.UserService
	refreshTokenService *app.RefreshTokenService
	userValidator       *validators.UserValidator
	userLoginValidator  *validators.UserLoginValidator
	userUpdateValidator *validators.UserUpdateValidator
}

func NewUserController(u *app.UserService, rt *app.RefreshTokenService) *UserController {
	return &UserController{
		userService:         u,
		refreshTokenService: rt,
		userValidator:       validators.NewUserValidator(),
		userLoginValidator:  validators.NewUserLoginValidator(),
		userUpdateValidator: validators.NewUserUpdateValidator(),
	}
}

func (c *UserController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := c.userValidator.ValidateAndMap(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		sevedUser, err := (*c.userService).Save(user)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}
		success(w, resources.MapDomainToUserDto(sevedUser))
	}
}

func (c *UserController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		autHeader := r.Header.Get("Authorization")
		token := autHeader[len(BEARER_SCHEMA):]

		params, err := parseUrlQuery(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}
		user, err := (*c.refreshTokenService).VerifyAccessToken(token)
		if err != nil {
			log.Print(err)
			return
		}
		usr, err := (*c.userService).FindOne(user.UserId, params)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}
		success(w, resources.MapDomainToUserDto(usr))
	}
}

func (c *UserController) PaginateAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseUrlQuery(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		users, err := (*c.userService).FindAll(params)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}
		success(w, resources.MapDomainToUserDtoCollection(users))
	}
}

func (c *UserController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := c.userUpdateValidator.ValidateAndMap(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		user.Id = id
		updateUser, err := (*c.userService).Update(user)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}
		success(w, resources.MapDomainToUserDto(updateUser))
	}
}

func (c *UserController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		err = (*c.userService).Delete(id)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}
		ok(w)
	}
}

func (c *UserController) LogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//validate input
		user, err := c.userValidator.ValidateAndMap(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}
		//login user
		userStored, err := (*c.userService).LogIn(user)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		storedToken, err := (*c.refreshTokenService).CreateRefreshToken(userStored)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		//set access token into the header
		accessToken, err := (*c.refreshTokenService).CreateAccessToken(storedToken)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		w.Header().Set("Authorization", accessToken)
		//generate success response
		success(w, resources.MapDomainTokenDto(accessToken))
	}
}

func (c *UserController) LogOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get refresh token
		autHeader := r.Header.Get("Authorization")
		token := autHeader[len(BEARER_SCHEMA):]

		user, err := (*c.refreshTokenService).VerifyAccessToken(token)
		if err != nil {
			log.Print(err)
			return
		}

		//delete refresh token
		err = (*c.refreshTokenService).DeleteSessionToken(user.UserId, user.Id)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		ok(w)
	}
}

func (c *UserController) CheckAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok(w)
	}
}
