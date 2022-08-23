package main

import (
	"context"
	"fmt"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"startUp/config"
	"startUp/internal/app"
	"startUp/internal/infra/database"
	"startUp/internal/infra/http/middlewares"
	"syscall"

	"startUp/internal/infra/http"
	"startUp/internal/infra/http/controllers"
)

func main() {
	exitCode := 0
	ctx, cancel := context.WithCancel(context.Background())

	//Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The sistem panicked!: %v\n", r)
			fmt.Printf("Stazk trace form panic: %s\n", string(debug.Stack()))
			exitCode = 1
		}
		os.Exit(exitCode)
	}()

	//Signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		fmt.Printf("Received signal '%s', stopping... \n", sig.String())
		cancel()
		fmt.Printf("sent cancel to all threads...")
	}()

	var conf = config.GetConfiguration()

	err := database.Migrate(conf)
	if err != nil {
		log.Fatalf("Unable to apply migration: %q\n", err)
	}

	ses, err := postgresql.Open(
		postgresql.ConnectionURL{
			User:     conf.DatabaseUser,
			Host:     conf.DatabaseHost,
			Password: conf.DatabasePassword,
			Database: conf.DatabaseName,
		})
	if err != nil {
		log.Fatalf("Unable to create new DB session: %q\n", err)
	}
	defer func(ses db.Session) {
		err = ses.Close()
		if err != nil {

		}
	}(ses)

	_, err = os.Stat(conf.FileStorageLocation)
	if err != nil {
		err = os.Mkdir(conf.FileStorageLocation, os.ModePerm)
	}
	if err != nil {
		log.Fatalf("Storage folder is not available %s", err)
	}

	//Coordinate
	coordinateRepository := database.NewRepository(&ses)
	coordinateService := app.NewService(&coordinateRepository)
	coordinateController := controllers.NewEventController(&coordinateService)

	//user
	userRepository := database.NewUserRepository(&ses)
	refreshTokenRepository := database.NewRefreshTokenRepository(&ses)
	userService := app.NewUserService(&userRepository)
	refreshTokenService := app.NewRefreshTokenService(&refreshTokenRepository, []byte(conf.AuthAccessKeySecret))
	userController := controllers.NewUserController(&userService, &refreshTokenService)

	authMiddleware := middlewares.AuthMiddleware(refreshTokenService)

	//HTTP Server
	err = http.Server(
		ctx,
		http.Router(
			authMiddleware,
			userController,
			coordinateController,
		),
	)

	if err != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}
}
