package resources

import "startUp/internal/domain/coordinates"

type CoordinateDTO struct {
	Id      int     `json:"id" db:"id"`
	MT      int     `json:"mt" db:"mt"`
	Axis    string  `json:"axis" db:"axis"`
	Horizon string  `json:"horizon" db:"horizon"`
	X       float64 `json:"x" db:"x"`
	Y       float64 `json:"y" db:"y"`
}

func MapDomainToCoordinateDTO(coordinate *coordinate.Coordinate) *CoordinateDTO {
	return &CoordinateDTO{
		Id:      coordinate.Id,
		MT:      coordinate.MT,
		Axis:    coordinate.Axis,
		Horizon: coordinate.Horizon,
		X:       coordinate.X,
		Y:       coordinate.Y,
	}
}

func MapDomainCoordinateCollection(coordinate []coordinate.Coordinate) *[]CoordinateDTO {
	var result []CoordinateDTO
	for _, t := range coordinate {
		dto := MapDomainToCoordinateDTO(&t)
		result = append(result, *dto)
	}
	return &result
}
