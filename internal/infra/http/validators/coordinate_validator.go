package validators

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"log"
	"net/http"
	"startUp/internal/domain/coordinates"
)

type CoordinateValidator struct {
	validator *validator.Validate
}

func NewCoordinateValidator() *CoordinateValidator {
	return &CoordinateValidator{
		validator: validator.New(),
	}
}

func (t CoordinateValidator) ValidateAndMap(request *http.Request) (*coordinate.Coordinate, error) {
	var coordinateReq coordinateRequest
	err := json.NewDecoder(request.Body).Decode(&coordinateReq)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return mapCoordinateRequestDomain(&coordinateReq), nil
}
