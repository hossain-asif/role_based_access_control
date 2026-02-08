package router

import (
	"go_project_structure/controllers"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", controllers.PingHandler)

	return router
}

