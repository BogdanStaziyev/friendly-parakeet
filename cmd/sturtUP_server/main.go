package main

import (
	"context"
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"startUp/config"
	"syscall"

	"startUp/internal/domain/coordinates"
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
	defer ses.Close()

	//Coordinate
	coordinateRepository := coordinate.NewRepository(&ses)
	coordinateService := coordinate.NewService(&coordinateRepository)
	coordinateController := controllers.NewEventController(&coordinateService)

	//HTTP Server
	err = http.Server(
		ctx,
		http.Router(
			coordinateController,
		),
	)

	if err != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}
}
