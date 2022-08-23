package controllers_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"startUp/internal/domain"
)

var userControllerTests = []*requestTest{
	{
		"Get paginated list",
		func(req *http.Request, migrator *migrate.Migrate) {
			//resetDB(migrator)
			entitiesToDb := 7
			userMocker(entitiesToDb)
			HeaderTokenMock(req, 2, 1, domain.ROLE_USER)
		},
		"/api/v1/user",
		"GET",
		``,
		http.StatusOK,
		`\[(?:{"id":\d+,"name":"User \d+","email":"Email\d+@email.com","role_id":\d+,"created_date":".{0,50}","updated_date":".{0,50}","deleted_date":".{0,50}"},?){7}\]`,
		"wrong list request",
	},
	{
		name: "LogIn ",
		init: func(req *http.Request, migrator *migrate.Migrate) {
			HeaderTokenMock(req, 1, 1, domain.ROLE_ADMIN)
		},
		url:    "/api/v1/user/login",
		method: "POST",
		bodyData: `{"email":"Email1@email.com",
    				"password":"password1"}`,
		expectedCode:   http.StatusOK,
		responseRegexg: `{"token":".{0,200}"}`,
		msg:            "wrong single task response",
	},
	{
		name: "Create object ",
		init: func(req *http.Request, migrator *migrate.Migrate) {
			HeaderTokenMock(req, 1, 1, domain.ROLE_ADMIN)
		},
		url:    "/api/v1/user",
		method: "POST",
		bodyData: `{"name":"User 8",
					"email":"Email8@email.com",
					"password":"password",
					"role_id": 2}`,
		expectedCode:   http.StatusOK,
		responseRegexg: `{"id":8,"name":"User 8","email":"Email8@email.com","role_id":2,"created_date":".{0,50}","updated_date":".{0,50}","deleted_date":".{0,50}"}`,
		msg:            "wrong single task response",
	},
	{
		"Get single object by token",
		func(req *http.Request, migrator *migrate.Migrate) {
			HeaderTokenMock(req, 8, 1, domain.ROLE_USER)
		},
		"/api/v1/user/profile",
		"GET",
		``,
		http.StatusOK,
		`{"id":8,"name":"User 8","email":"Email8@email.com","role_id":2,"created_date":".{0,50}","updated_date":".{0,50}","deleted_date":".{0,50}"}`,
		"wrong single task response",
	},

	{
		"Update object by Id ",
		func(req *http.Request, migrator *migrate.Migrate) {
			HeaderTokenMock(req, 1, 1, domain.ROLE_ADMIN)
		},
		"/api/v1/user/8",
		"PUT",
		`{	"name":"name2",
					"email":"Email8@email.com",
					"role_id": 2}`,
		http.StatusOK,
		`{"id":\d+,"name":"name2","email":"Email8@email.com","role_id":2,"created_date":".{0,50}","updated_date":".{0,50}","deleted_date":".{0,50}"}`,
		"wrong single task response",
	},
	{
		"Delete object by Id ",
		func(req *http.Request, migrator *migrate.Migrate) {
			HeaderTokenMock(req, 1, 1, domain.ROLE_ADMIN)
		},
		"/api/v1/user/7",
		"DELETE",
		``,
		http.StatusOK,
		``,
		"wrong single task response",
	},
	{
		"Check Auth",
		func(req *http.Request, migrator *migrate.Migrate) {
			HeaderTokenMock(req, 1, 1, domain.ROLE_ADMIN)
		},
		"/api/v1/user/checkauth",
		"GET",
		``,
		http.StatusOK,
		``,
		"wrong single task response",
	},
	{
		"Check Auth",
		func(req *http.Request, migrator *migrate.Migrate) {
			HeaderTokenMock(req, 1, 1, domain.ROLE_ADMIN)
		},
		"/api/v1/user/checkauth",
		"GET",
		``,
		http.StatusOK,
		``,
		"wrong single task response",
	},
	{
		"LogOut",
		func(req *http.Request, migrator *migrate.Migrate) {
			HeaderTokenMock(req, 1, 1, domain.ROLE_ADMIN)
		},
		"/api/v1/user/logout",
		"POST",
		``,
		http.StatusOK,
		``,
		"wrong single task response",
	},
}

func userMocker(n int) []domain.User {
	cts := make([]domain.User, 0, n)
	userModel := domain.User{
		Name:     fmt.Sprintf("User %d", 1),
		Email:    fmt.Sprintf("Email%d@email.com", 1),
		Password: fmt.Sprintf("password%d", 1),
		Role:     1,
	}
	ct, err := userService.Save(&userModel)
	if err != nil {
		log.Printf("userMocker() failed: %s", err)

	}
	cts = append(cts, *ct)

	for i := 2; i <= n; i++ {
		userModel = domain.User{
			Name:     fmt.Sprintf("User %d", i),
			Email:    fmt.Sprintf("Email%d@email.com", i),
			Password: fmt.Sprintf("password%d", i),
			Role:     2,
		}
		ct, err = userService.Save(&userModel)
		if err != nil {
			log.Printf("userMocker() failed: %s", err)

		}
		cts = append(cts, *ct)
	}
	return cts
}
