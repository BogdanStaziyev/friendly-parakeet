package validators

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"log"
	"net/http"
	"startUp/internal/domain"
)

type CoordinateValidator struct {
	validator *validator.Validate
}

func NewCoordinateValidator() *CoordinateValidator {
	return &CoordinateValidator{
		validator: validator.New(),
	}
}

func (t CoordinateValidator) ValidateAndMap(request *http.Request) (*domain.Coordinate, error) {
	var coordinateReq coordinateRequest
	err := json.NewDecoder(request.Body).Decode(&coordinateReq)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = t.validator.Struct(coordinateReq)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return mapCoordinateRequestDomain(&coordinateReq), nil
}
