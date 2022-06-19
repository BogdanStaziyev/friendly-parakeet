package controllers

import (
	"fmt"
	"net/http"
	"startUp/internal/domain/event"
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
