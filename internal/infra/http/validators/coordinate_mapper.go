package validators

import (
	"startUp/internal/domain"
)

type coordinateRequest struct {
	MT      int64   `json:"mt" validate:"required"`
	Axis    string  `json:"axis" validate:"required"`
	Horizon string  `json:"horizon" validate:"required"`
	X       float64 `json:"x" validate:"required"`
	Y       float64 `json:"y" validate:"required"`
}

func mapCoordinateRequestDomain(coordinateRequest *coordinateRequest) *domain.Coordinate {
	return &domain.Coordinate{
		MT:      coordinateRequest.MT,
		Axis:    coordinateRequest.Axis,
		Horizon: coordinateRequest.Horizon,
		X:       coordinateRequest.X,
		Y:       coordinateRequest.Y,
	}
}
