package controllers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"startUp/internal/domain/coordinates"
	"startUp/internal/domain/resources"
	"startUp/internal/infra/http/validators"
	"strconv"
)

type EventController struct {
	service   *coordinate.Service
	validator *validators.CoordinateValidator
}

func NewEventController(s *coordinate.Service) *EventController {
	return &EventController{
		service:   s,
		validator: validators.NewCoordinateValidator(),
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coordinates, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindeAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}
		err = success(w, resources.MapDomainCoordinateCollection(coordinates))
		if err != nil {
			fmt.Printf("EventController.FindAll: %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(request, "id"), 10, 64)
		if err != nil {
			fmt.Printf("CoordinateController.FindeOne(): %s", err)
			err = internalServerError(writer, err)
			if err != nil {
				fmt.Printf("CoordinateController.FindeOne(): %s", err)
			}
			return
		}
		coordinates, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("CoordinateController.FindeOne(): %s", err)
			err = internalServerError(writer, err)
			if err != nil {
				fmt.Printf("CoordinateController.FindeOne(): %s", err)
			}
			return
		}

		err = success(writer, resources.MapDomainToCoordinateDTO(coordinates))
		if err != nil {
			fmt.Printf("CoordinateController.FindeOne(): %s", err)
		}
	}
}

func (c *EventController) AddCoordinate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		coordinates, err := c.validator.ValidateAndMap(request)
		if err != nil {
			log.Print(writer, err)
			badRequest(writer, err)
			return
		}

		createCoordinate, err := (*c.service).AddCoordinate(coordinates)
		if err != nil {
			internalServerError(writer, err)
			return
		}

		err = success(writer, resources.MapDomainToCoordinateDTO(createCoordinate))
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *EventController) UpdateCoordinate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		coordinates, err := c.validator.ValidateAndMap(request)
		if err != nil {
			log.Println(writer, err)
			badRequest(writer, err)
			return
		}
		err = (*c.service).UpdateCoordinate(coordinates)
		if err != nil {
			log.Println(writer, err)
			internalServerError(writer, err)
			return
		}
		ok(writer)
	}
}

func (c *EventController) InverseTask() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
