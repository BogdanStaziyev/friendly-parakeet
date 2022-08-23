package controllers_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/stretchr/testify/require"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"startUp/config"
	"startUp/internal/app"
	"startUp/internal/domain"
	"startUp/internal/infra/database"
	"startUp/internal/infra/http/controllers"
	"startUp/internal/infra/http/middlewares"
	"testing"

	userHttp "startUp/internal/infra/http"
)

var (
	refreshTokenService app.RefreshTokenService
	userService         app.UserService
	coordinateService   app.Service
)

type requestTest struct {
	name           string
	init           func(*http.Request, *migrate.Migrate)
	url            string
	method         string
	bodyData       string
	expectedCode   int
	responseRegexg string
	msg            string
}

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestController(t *testing.T) {

	var conf = &config.Configuration{
		DatabaseName:        "postgres",
		DatabaseHost:        "localhost:54322",
		DatabaseUser:        "postgres",
		DatabasePassword:    "password",
		MigrateToVersion:    "latest",
		MigrationLocation:   "../../database/migrations",
		FileStorageLocation: "",
	}

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		conf.DatabaseUser,
		conf.DatabasePassword,
		conf.DatabaseHost,
		conf.DatabaseName,
	)

	internalSqlDriver, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Unable to open DB connection: %q\n", err)
	}

	driver, err := postgres.WithInstance(internalSqlDriver, &postgres.Config{})
	migrator, err := migrate.NewWithDatabaseInstance(
		"file://"+conf.MigrationLocation,
		conf.DatabaseName,
		driver)
	if err != nil {
		log.Fatalf("Unable to create Migrator: %q\n", err)
	}
	sess, err := postgresql.New(internalSqlDriver)
	if err != nil {
		log.Fatalf("Uneble to create Migrator: %q\n", err)
	}

	refreshTokenRepository := database.NewRefreshTokenRepository(&sess)
	refreshTokenService = app.NewRefreshTokenService(&refreshTokenRepository, []byte(conf.AuthAccessKeySecret))
	authMiddleware := middlewares.AuthMiddleware(refreshTokenService)

	userRepository := database.NewUserRepository(&sess)
	userService = app.NewUserService(&userRepository)
	userController := controllers.NewUserController(&userService, &refreshTokenService)

	coordinateRepository := database.NewRepository(&sess)
	coordinateService = app.NewService(&coordinateRepository)
	coordinateController := controllers.NewEventController(&coordinateService)

	router := userHttp.Router(
		authMiddleware,
		userController,
		coordinateController,
	)
	//Reset DB ti clean state
	resetDB(migrator)
	//Truncate from tables
	prepareTestDB(sess)

	invertOverTest(t, "userControllerTests", userControllerTests, router, migrator)
	invertOverTest(t, "CoordinatesControllerTest", coordinateControllerTest, router, migrator)
}

func invertOverTest(t *testing.T, name string, tests []*requestTest, router http.Handler, migrator *migrate.Migrate) {
	for _, tt := range tests {
		t.Run(name+tt.name, func(t *testing.T) {
			fmt.Printf("[%-6s] %-35s", tt.method, tt.url)
			bodyData := tt.bodyData
			req, err := http.NewRequest(tt.method, tt.url, bytes.NewBufferString(bodyData))
			require.NoError(t, err)
			req.Header.Set("Content-type", "application/json")

			tt.init(req, migrator)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			body := w.Body.String()

			fmt.Printf("[%d]\n", w.Code)
			require.Equal(t, tt.expectedCode, w.Code, "Response Status - "+tt.msg+"\nBody:\n"+body)
			require.Regexp(t, tt.responseRegexg, body, "Response content - "+tt.msg)
		})
	}
}

//func HeaderTokenMock()
func HeaderTokenMock(req *http.Request, u, id int64, role domain.Role) {
	accessToken, _ := refreshTokenService.CreateAccessToken(&domain.RefreshToken{
		Id:       id,
		UserId:   u,
		UserRole: role,
	})
	bearer := "Bearer " + accessToken
	req.Header.Set("Authorization", bearer)
}

func prepareTestDB(sess db.Session) {
	err := sess.Collection(database.UserTable).Truncate()
	if err != nil {
		log.Print("prepareTestDB CoordinateTable error:", err)
	}
	err = sess.Collection(database.CoordinateTable).Truncate()
	if err != nil {
		log.Print("prepareTestDB CoordinateTable error:", err)
	}
}

func resetDB(migrator *migrate.Migrate) {
	err := migrator.Down()
	if err != nil {
		log.Printf("migrator down: %q\n", err)
	}
	err = migrator.Up()
	if err != nil {
		log.Printf("migrator up: %q\n", err)
	}
}
