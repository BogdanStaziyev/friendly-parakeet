package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"startUp/internal/domain/coordinates"
	"strconv"
)

type EventController struct {
	service *coordinate.Service
}

func NewEventController(s *coordinate.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coordinates, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindeAll(): %s", err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, coordinates)
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
		err = success(writer, coordinates)
		if err != nil {
			fmt.Printf("CoordinateController.FindeOne(): %s", err)
		}
	}
}

func (c *EventController) AddCoordinate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var coordinate coordinate.Coordinate
		json.NewDecoder(request.Body).Decode(&coordinate)
		if coordinate.X == 0 {
			log.Fatal("X not exist")
		} else if coordinate.Y == 0 {
			log.Fatal("Y not exist")
		} else if coordinate.MT == 0 {
			fmt.Println("MT not exist")
			return
		}
		err := (*c.service).AddCoordinate(coordinate)
		if err != nil {
			fmt.Printf("CoordinateController.AddCoordinate(): %s", err)
			err = internalServerError(writer, err)
			if err != nil {
				fmt.Printf("CoordinateController.AddCoordinate(): %s", err)
			}
		}
	}
}

func (c *EventController) UpdateCoordinate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func (c *EventController) InverseTask() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
