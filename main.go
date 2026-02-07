package main

import "go_project_structure/app"

func main() {

	cfg := app.Config{
		Addr: ":3000",
	}

	app := app.Application{
		Config: cfg,
	}

	app.Run()
}

