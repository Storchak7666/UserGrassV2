package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/test_server/internal/infra/http/controllers"
)

func Router(userController *controllers.UserController) http.Handler {
	router := chi.NewRouter()

	// Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Use(middleware.RedirectSlashes)

		healthRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})
	})

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes)

		apiRouter.Route("/grass", func(apiRouter chi.Router) {
			AddUserRoutes(&apiRouter, userController)

			apiRouter.Handle("/*", NotFoundJSON())
		})
	})

	return router
}

func AddUserRoutes(router *chi.Router, userController *controllers.UserController) {
	(*router).Route("/users", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/FindAll",
			userController.FindAll(),
		)
		apiRouter.Get(
			"/FindById/{id}",
			userController.FindById(),
		)

		apiRouter.Post(
			"/Create",
			userController.CreateUser(),
		)
		apiRouter.Put(
			"/UpdateById",
			userController.UpdateById(),
		)
	})
}
