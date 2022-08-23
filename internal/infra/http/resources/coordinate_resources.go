package resources

import (
	"startUp/internal/domain"
	"time"
)

type CoordinateDTO struct {
	Id          int64     `json:"id"`
	MT          int64     `json:"mt"`
	Axis        string    `json:"axis"`
	Horizon     string    `json:"horizon"`
	X           float64   `json:"x"`
	Y           float64   `json:"y"`
	UserId      int64     `db:"user_id"`
	CreatedDate time.Time `db:"created_date"`
	UpdatedDate time.Time `db:"updated_date"`
	DeletedDate time.Time `db:"deleted_date"`
}

func MapDomainToCoordinateDTO(coordinate *domain.Coordinate) *CoordinateDTO {
	return &CoordinateDTO{
		Id:          coordinate.Id,
		MT:          coordinate.MT,
		Axis:        coordinate.Axis,
		Horizon:     coordinate.Horizon,
		X:           coordinate.X,
		Y:           coordinate.Y,
		UserId:      coordinate.UserID,
		CreatedDate: coordinate.CreatedDate,
		UpdatedDate: coordinate.UpdatedDate,
		DeletedDate: coordinate.DeletedDate,
	}
}

func MapDomainCoordinateCollection(coordinate []domain.Coordinate) *[]CoordinateDTO {
	var result []CoordinateDTO
	for _, t := range coordinate {
		dto := MapDomainToCoordinateDTO(&t)
		result = append(result, *dto)
	}
	return &result
}
