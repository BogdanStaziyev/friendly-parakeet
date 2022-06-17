package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"startUp/internal/domain/event"
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

	//Event
	coordinateRepository := coordinate.NewRepository()
	coordinateService := coordinate.NewService(&coordinateRepository)
	coordinateController := controllers.NewEventController(&coordinateService)

	//HTTP Server
	err := http.Server(
		ctx,
		http.Router(
			coordinateController,
		),
	)

	if err != nil{
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}
}