package validators

import "startUp/internal/domain/coordinates"

type coordinateRequest struct {
	Id      int     `json:"id"`
	MT      int     `json:"mt" validate:"required"`
	Axis    string  `json:"axis" validate:"required"`
	Horizon string  `json:"horizon" validate:"required"`
	X       float64 `json:"x" validate:"required"`
	Y       float64 `json:"y" validate:"required"`
}

func mapCoordinateRequestDomain(coordinateRequest *coordinateRequest) *coordinate.Coordinate {
	return &coordinate.Coordinate{
		Id:      coordinateRequest.Id,
		MT:      coordinateRequest.MT,
		Axis:    coordinateRequest.Axis,
		Horizon: coordinateRequest.Horizon,
		X:       coordinateRequest.X,
		Y:       coordinateRequest.Y,
	}
}
