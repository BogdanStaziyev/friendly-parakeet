package controllers_test

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"net/http"
	"startUp/internal/domain"
)

var coordinateControllerTest = []*requestTest{
	{
		"Get all list ",
		func(r *http.Request, migrator *migrate.Migrate) {
			resetDB(migrator)
			entitiesToDB := 10
			coordinateMocker(entitiesToDB)
		},
		"/v1/coordinates",
		"GET",
		``,
		http.StatusOK,
		`\[(?:{"id":\d+,"mt":\d+,"axis":"coordinate \d+","horizon":"coordinate \d+","x":\d+(?:[\.,]\d+)+,"y":\d+(?:[\.,]\d+)?},?){10}\]`,
		"wrong list request",
	},
	{
		"Get single object ID=5 ",
		func(req *http.Request, migrator *migrate.Migrate) {},
		"/v1/coordinates/5",
		"GET",
		``,
		http.StatusOK,
		`{"id":\d+,"mt":\d+,"axis":"coordinate \d+","horizon":"coordinate \d+","x":\d+(?:[\.,]\d+)+,"y":\d+(?:[\.,]\d+)}`,
		"wrong single task response",
	},
	{
		"Create object ",
		func(req *http.Request, migrator *migrate.Migrate) {},
		"/v1/coordinates/add",
		"POST",
		`{
			"mt": 34,
			"axis": "created object",
			"horizon": "created object",
			"x": 333333.333,
			"y": 222222.222
		}`,
		http.StatusOK,
		`{"id":\d+,"mt":34,"axis":"created object","horizon":"created object","x":333333.333,"y":222222.222}`,
		"wrong single task response",
	},
	{
		"Update object ID=7 ",
		func(req *http.Request, migrator *migrate.Migrate) {},
		"/v1/coordinates/update",
		"PUT",
		`{
			"id": 7,
			"mt": 34,
			"axis": "created object",
			"horizon": "created object",
			"x": 333333.333,
			"y": 222222.222
		}`,
		http.StatusOK,
		``,
		"wrong single task response",
	},
	{
		"Delete object by ID=2",
		func(r *http.Request, migrate *migrate.Migrate) {},
		"/v1/coordinates/2",
		"DELETE",
		``,
		http.StatusOK,
		``,
		"wrong single object",
	},
	{
		"Get single object by Deleted ID=2",
		func(req *http.Request, migrator *migrate.Migrate) {},
		"/v1/coordinates/2",
		`GET`,
		``,
		http.StatusNotFound,
		`service FindOne: coordinateRepository FindOne: upper: no more rows in this result set`,
		"wrong single task response",
	},
	{
		"Create object 111",
		func(req *http.Request, migrator *migrate.Migrate) {},
		"/v1/coordinates/add",
		"POST",
		`{
			"mt": 111,
			"axis": "188",
			"horizon": "1340",
			"x": 25994.292,
			"y": 35476.775
		}`,
		http.StatusOK,
		`{"id":\d+,"mt":111,"axis":"188","horizon":"1340","x":25994.292,"y":35476.775}`,
		"wrong single task response",
	},
	{
		"Create object 109",
		func(req *http.Request, migrator *migrate.Migrate) {},
		"/v1/coordinates/add",
		"POST",
		`{
			"mt": 109,
			"axis": "189",
			"horizon": "1340",
			"x": 25969.153,
			"y": 35462.677
		}`,
		http.StatusOK,
		`{"id":\d+,"mt":109,"axis":"189","horizon":"1340","x":25969.153,"y":35462.677}`,
		"wrong single task response",
	},
	{
		"Get result real coordinates MT 111-109",
		func(req *http.Request, migrator *migrate.Migrate) {},
		"/v1/coordinates/13/12",
		`GET`,
		``,
		http.StatusOK,
		`"Результат обчислення зворотньої геодезичної задачі: 29° 17′ 1″ "`,
		"wrong invert result response",
	},
}

func coordinateMocker(n int) []domain.Coordinate {
	coordinates := make([]domain.Coordinate, 0, n)
	for i := 1; i <= n; i++ {
		coordinateModel := domain.Coordinate{
			MT:      int64(i),
			Axis:    fmt.Sprintf("coordinate %d", i),
			Horizon: fmt.Sprintf("coordinate %d", i),
			X:       25252.252 + float64(i),
			Y:       35353.352 + float64(i),
		}
		coordinete, err := coordinateService.AddCoordinate(&coordinateModel)
		if err != nil {
			log.Printf("CoordinateMocker() dailed: %s", err)
		}
		coordinates = append(coordinates, *coordinete)
	}
	return coordinates
}
