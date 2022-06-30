package http

import (
	"net/http"
	"startUp/internal/infra/http/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(eventController *controllers.CoordinateController) http.Handler {
	router := chi.NewRouter()

	//Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})
	})

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes)

		apiRouter.Route("/v1", func(apiRouter chi.Router) {

			apiRouter.Group(func(apiRouter chi.Router) {
				AddEventRoutes(&apiRouter, eventController)
				apiRouter.Handle("/*", NotFoundJSON())
			})
			apiRouter.Handle("/*", NotFoundJSON())
		})
	})
	return router
}

func AddEventRoutes(router *chi.Router, coordinateController *controllers.CoordinateController) {
	(*router).Route("/coordinates", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			coordinateController.FindAll(),
		)
		apiRouter.Get(
			"/{id}",
			coordinateController.FindOne(),
		)
		apiRouter.Post(
			"/add",
			coordinateController.AddCoordinate(),
		)
		apiRouter.Put(
			"/update",
			coordinateController.UpdateCoordinate(),
		)
		apiRouter.Delete(
			"/{id}",
			coordinateController.DeleteCoordinate(),
		)
		apiRouter.Get(
			"/{firstId}/{secondId}",
			coordinateController.InverseTask(),
		)
	})
}
