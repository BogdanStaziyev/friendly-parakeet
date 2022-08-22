package http

import (
	"github.com/go-chi/cors"
	"net/http"
	"startUp/internal/infra/http/controllers"
	"startUp/internal/infra/http/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HandlerFuncWrapper func(http.HandlerFunc) http.HandlerFunc
type HandlerMiddleware func(http.Handler) http.Handler

func Router(
	userController *controllers.UserController,
	authMiddleware HandlerMiddleware,
	coordinateController *controllers.CoordinateController,
	//TODO: uncomment to turn on commandController
	//commandController *controllers.CommandController,
) http.Handler {

	router := chi.NewRouter()

	// Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Use(middleware.RedirectSlashes)

		healthRouter.Route("/api/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})
	})

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes, cors.Handler(cors.Options{
			AllowedOrigins: []string{
				"https://*",
				"http://*",
			},
			AllowedMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"OPTIONS",
			},
			AllowedHeaders: []string{
				"Accept",
				"Authorization",
				"Content-Type",
				"X-CSRF-Token",
			},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

		apiRouter.Route("/api/v1", func(apiRouter chi.Router) {
			apiRouter.Group(func(apiRouter chi.Router) {
				//TODO: uncoment to turn on authorization cheks
				apiRouter.Use(authMiddleware)
				//CommandRouter(&apiRouter, commandController)
				CoordinateRouter(&apiRouter, coordinateController)
				UserRouter(&apiRouter, userController)
				apiRouter.Handle("/*", NotFoundJSON())

			})
			apiRouter.Post(
				"/user/login",
				userController.LogIn(),
			)
			apiRouter.Handle("/*", NotFoundJSON())
		})
	})

	return router
}

func UserRouter(router *chi.Router, userController *controllers.UserController) {
	(*router).Route("/user", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			userController.PaginateAll(),
		)
		apiRouter.Get(
			"/profile",
			userController.FindOne(),
		)
		apiRouter.Get(
			"/checkauth",
			userController.CheckAuth(),
		)
		apiRouter.Post(
			"/logout",
			userController.LogOut(),
		)
		apiRouterAdminOnly := apiRouter.With(middlewares.AdminOnli)
		apiRouterAdminOnly.Post(
			"/",
			userController.Save(),
		)
		apiRouterAdminOnly.Put(
			"/{id}",
			userController.Update(),
		)
		apiRouterAdminOnly.Delete(
			"/{id}",
			userController.Delete(),
		)
	})
}

func CoordinateRouter(router *chi.Router, coordinateController *controllers.CoordinateController) {
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
