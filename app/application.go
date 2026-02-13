package app

import (
	"fmt"
	dbConfig "go_project_structure/config/db"
	config "go_project_structure/config/env"
	"go_project_structure/internal/user"
	"go_project_structure/router"

	"net/http"
	"time"
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

	db, err := dbConfig.SetupDB()
	if err != nil {
		fmt.Println("Error setting up database.")
		return err
	}

	ur := user.NewUserRepository(db)
	us := user.NewUserService(ur)
	uc := user.NewUserController(us)
	uRouter := router.NewUserRouter(uc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(uRouter),
		ReadTimeout:  10 * time.Second, // Set read timeout to 10 seconds
		WriteTimeout: 10 * time.Second, // Set write timeout to 10 seconds
	}

	fmt.Println("Starting server on port", app.Config.Addr)

	return server.ListenAndServe()
}
