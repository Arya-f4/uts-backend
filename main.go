
package main

import (
	"golang-train/config"
	"golang-train/database"
	"golang-train/router"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	dbPool := database.NewPostgresConnection(cfg)
	defer dbPool.Close()

	app := config.NewFiber()

	router.SetupRoutes(app, dbPool, cfg)

	log.Printf("Server is running on port %s", cfg.ServerPort)
	err = app.Listen(":" + cfg.ServerPort)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}


