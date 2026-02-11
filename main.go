package main

import (
	"go_project_structure/app"
	config "go_project_structure/config/env"
	db "go_project_structure/config/db"
)

func main() {
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)

	db.SetupDB()
	app.Run()
}
