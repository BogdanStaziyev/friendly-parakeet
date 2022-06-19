package coordinate

type Coordinate struct {
	Id      int     `json:"id" db:"id"`
	MT      int     `json:"mt" db:"mt"`
	Axis    string  `json:"axis" db:"axis"`
	Horizon string  `json:"horizon" db:"horizon"`
	X       float64 `json:"x" db:"x"`
	Y       float64 `json:"y" db:"y"`
}
