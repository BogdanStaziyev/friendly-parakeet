package validators

import (
	"startUp/internal/domain"
)

type coordinateRequest struct {
	ID      int64   `json:"id"`
	MT      int64   `json:"mt" validate:"required"`
	Axis    string  `json:"axis" validate:"required"`
	Horizon string  `json:"horizon" validate:"required"`
	X       float64 `json:"x" validate:"required"`
	Y       float64 `json:"y" validate:"required"`
}

func mapCoordinateRequestDomain(coordinateRequest *coordinateRequest) *domain.Coordinate {
	return &domain.Coordinate{
		Id:      coordinateRequest.ID,
		MT:      coordinateRequest.MT,
		Axis:    coordinateRequest.Axis,
		Horizon: coordinateRequest.Horizon,
		X:       coordinateRequest.X,
		Y:       coordinateRequest.Y,
	}
}
