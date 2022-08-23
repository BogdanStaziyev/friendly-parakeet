package domain

import "time"

type Coordinate struct {
	Id          int64
	MT          int64
	Axis        string
	Horizon     string
	X           float64
	Y           float64
	UserID      int64
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate time.Time
}
