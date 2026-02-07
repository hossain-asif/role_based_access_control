package main

import "go_project_structure/app"

func main() {

	cfg := app.NewConfig(":3000")
	app := app.NewApplication(cfg)

	app.Run()
}

