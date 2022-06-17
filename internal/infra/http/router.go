package http

import (
	"net/http"
	"startUp/internal/infra/http/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(eventController *controllers.EventController) http.Handler  {
	router := chi.NewRouter()

	//Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})
	})
	
	router.Group(func(apiRouter chi.Router){
		apiRouter.Use(middleware.RedirectSlashes)
		
		apiRouter.Route("/v1", func(apiRouter chi.Router){
			
			apiRouter.Group(func(apiRouter chi.Router) {
				AddEventRoutes(&apiRouter, eventController)
				apiRouter.Handle("/*", NotFoundJSON())
			})
			apiRouter.Handle("/*", NotFoundJSON())
		})
	})
		return router
}

func AddEventRoutes(router *chi.Router, eventController *controllers.EventController) {
	(*router).Route("/coordinates", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			eventController.FindAll(),
			)
	})
}