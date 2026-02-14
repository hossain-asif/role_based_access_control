package app

import (
	"fmt"
	dbConfig "go_project_structure/config/db"
	config "go_project_structure/config/env"
	"go_project_structure/internal/router"

	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// Config holds the configuration for the server.
type Config struct {
	Addr string // PORT
}

// constructor for Config
func NewConfig() Config {
	port := config.GetString("PORT", ":8080")
	return Config{
		Addr: port,
	}
}

type Application struct {
	Config Config
}

// constructor for Application
func NewApplication(config Config) Application {
	return Application{
		Config: config,
	}
}

func (app *Application) Run() error {

	rootRouter := chi.NewRouter()

	db, err := dbConfig.SetupDB()
	if err != nil {
		fmt.Println("Error setting up database.")
		return err
	}

	for _, registerFn := range router.DomainRegistries {
		registerFn(db, rootRouter)
	}

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      rootRouter,       
		ReadTimeout:  10 * time.Second, // Set read timeout to 10 seconds
		WriteTimeout: 10 * time.Second, // Set write timeout to 10 seconds
	}

	fmt.Println("Starting server on port", app.Config.Addr)

	return server.ListenAndServe()
}
