package resources

import "startUp/internal/domain/coordinates"

type CoordinateDTO struct {
	Id      int64   `json:"id"`
	MT      int64   `json:"mt"`
	Axis    string  `json:"axis"`
	Horizon string  `json:"horizon"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
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
