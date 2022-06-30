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

type CoordinateController struct {
	service   *coordinate.Service
	validator *validators.CoordinateValidator
}

func NewEventController(s *coordinate.Service) *CoordinateController {
	return &CoordinateController{
		service:   s,
		validator: validators.NewCoordinateValidator(),
	}
}

func (c *CoordinateController) AddCoordinate() http.HandlerFunc {
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

func (c *CoordinateController) UpdateCoordinate() http.HandlerFunc {
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

func (c *CoordinateController) DeleteCoordinate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(request, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.DeleteCoordinate(): %s", err)
			badRequest(writer, err)
			return
		}

		err = (*c.service).DeleteCoordinate(id)
		if err != nil {
			fmt.Printf("EventController.DeleteCoordinate(): %s", err)
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

func (c *CoordinateController) FindOne() http.HandlerFunc {
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

func (c *CoordinateController) InverseTask() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		firstId, err := strconv.ParseInt(chi.URLParam(request, "firstId"), 10, 64)
		if err != nil {
			fmt.Printf("CoordinateController.InverseFirst(): %s", err)
			err = internalServerError(writer, err)
			if err != nil {
				fmt.Printf("CoordinateController.InverseFirst(): %s", err)
			}
			return
		}
		secondId, err := strconv.ParseInt(chi.URLParam(request, "secondId"), 10, 64)
		if err != nil {
			fmt.Printf("CoordinateController.InverseSecond(): %s", err)
			err = internalServerError(writer, err)
			if err != nil {
				fmt.Printf("CoordinateController.InverseSecond(): %s", err)
			}
			return
		}
		coordinatAxis, err := (*c.service).InverseTask(firstId, secondId)
		if err != nil {
			fmt.Printf("CoordinateController.InverseFirst(): %s", err)
			err = internalServerError(writer, err)
			if err != nil {
				fmt.Printf("CoordinateController.InverseFirst(): %s", err)
			}
			return
		}
		err = success(writer, coordinatAxis)
		if err != nil {
			fmt.Printf("CoordinateController.InverseFirst(): %s", err)
		}
	}
}
