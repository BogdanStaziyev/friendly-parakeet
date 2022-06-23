package validators

import "startUp/internal/domain/coordinates"

type coordinateRequest struct {
	Id      int     `json:"id" db:"id"`
	MT      int     `json:"mt" db:"mt" validate:"required"`
	Axis    string  `json:"axis" db:"axis" validate:"required"`
	Horizon string  `json:"horizon" db:"horizon" validate:"required"`
	X       float64 `json:"x" db:"x" validate:"required"`
	Y       float64 `json:"y" db:"y" validate:"required"`
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
