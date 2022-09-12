package controllers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"startUp/internal/app"
	"startUp/internal/infra/http/resources"
	"startUp/internal/infra/http/validators"
	"strconv"
	"strings"
)

type CoordinateController struct {
	service             *app.CoordinateService
	validator           *validators.CoordinateValidator
	refreshTokenService *app.RefreshTokenService
}

func NewCoordinateController(s *app.CoordinateService, rt *app.RefreshTokenService) *CoordinateController {
	return &CoordinateController{
		service:             s,
		validator:           validators.NewCoordinateValidator(),
		refreshTokenService: rt,
	}
}

func (c *CoordinateController) AddCoordinate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")
		token := authHeader[len(BEARER_SCHEMA):]

		user, err := (*c.refreshTokenService).VerifyAccessToken(token)
		if err != nil {
			log.Println(writer, err)
		}
		coordinates, err := c.validator.ValidateAndMap(request)
		if err != nil {
			log.Print(writer, err)
			badRequest(writer, err)
			return
		}

		coordinates.UserID = user.UserId
		createCoordinate, err := (*c.service).AddCoordinate(coordinates)
		if err != nil {
			log.Println(err)
			internalServerError(writer, err)
			return
		}
		success(writer, resources.MapDomainToCoordinateDTO(createCoordinate))
	}
}

func (c *CoordinateController) UpdateCoordinate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")
		token := authHeader[len(BEARER_SCHEMA):]

		user, err := (*c.refreshTokenService).VerifyAccessToken(token)
		if err != nil {
			log.Println(writer, err)
		}
		coordinates, err := c.validator.ValidateAndMap(request)
		if err != nil {
			log.Println(err)
			badRequest(writer, err)
			return
		}

		coordinates.UserID = user.UserId
		err = (*c.service).UpdateCoordinate(coordinates)
		if err != nil {
			log.Println(err)
			internalServerError(writer, err)
			return
		}
		ok(writer)
	}
}

func (c *CoordinateController) DeleteCoordinate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(request, "id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(writer, fmt.Errorf("expected 'id' to be an integer, was given: %s", chi.URLParam(request, "id")))
			return
		}

		err = (*c.service).DeleteCoordinate(id)
		if err != nil {
			log.Println(err)
			internalServerError(writer, err)
			return
		}

		ok(writer)
	}
}

func (c *CoordinateController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coordinates, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindeAll(): %s", err)
			internalServerError(w, err)
			return
		}
		success(w, resources.MapDomainCoordinateCollection(coordinates))
	}
}

func (c *CoordinateController) FindOne() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(request, "id"), 10, 64)
		if err != nil {
			log.Println(err)
			badRequest(writer, fmt.Errorf("expected 'id' to be an integer, was given: %s", chi.URLParam(request, "id")))
			return
		}
		coordinates, err := (*c.service).FindOne(id)
		if err != nil {
			log.Println("coordinateService error", err)
			if strings.HasSuffix(err.Error(), "upper: no more rows in this result set") {
				notFound(writer, err)
			} else {
				notFound(writer, err)
			}
			return
		}
		success(writer, resources.MapDomainToCoordinateDTO(coordinates))
	}
}

func (c *CoordinateController) InverseTask() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		firstId, err := strconv.ParseInt(chi.URLParam(request, "firstId"), 10, 64)
		if err != nil {
			fmt.Printf("CoordinateController.InverseFirst(): %s", err)
			internalServerError(writer, err)
			if err != nil {
				fmt.Printf("CoordinateController.InverseFirst(): %s", err)
			}
			return
		}
		secondId, err := strconv.ParseInt(chi.URLParam(request, "secondId"), 10, 64)
		if err != nil {
			fmt.Printf("CoordinateController.InverseSecond(): %s", err)
			internalServerError(writer, err)
			if err != nil {
				fmt.Printf("CoordinateController.InverseSecond(): %s", err)
			}
			return
		}
		coordinateAxis, err, coordinateOne, coordinateTwo := (*c.service).InverseTask(firstId, secondId)
		if err != nil {
			fmt.Printf("CoordinateController.InverseFirst(): %s", err)
			internalServerError(writer, err)
			if err != nil {
				fmt.Printf("CoordinateController.InverseFirst(): %s", err)
			}
			return
		}
		success(writer, coordinateAxis)
		if err != nil {
			fmt.Printf("CoordinateController.InverseFirst(): %s", err)
		}
		success(writer, coordinateOne)
		if err != nil {
			fmt.Printf("CoordinateController.InverseFirst(): %s", err)
		}
		success(writer, coordinateTwo)
		if err != nil {
			fmt.Printf("CoordinateController.InverseFirst(): %s", err)
		}
	}
}
