package main

import (
	"go_project_structure/app"
	config "go_project_structure/config/env"
)

func main() {
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)

	app.Run()
}

// func main() {
// 	config.Load()
// 	cfg := app.NewConfig()
// 	app := app.NewApplication(cfg)

// 	go func() {
// 		if err := app.Run(); err != nil && err != http.ErrServerClosed {
// 			fmt.Printf("Server error: %v\n", err)
// 			os.Exit(1)
// 		}
// 	}()

// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit

// }
