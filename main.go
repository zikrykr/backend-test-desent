package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zikrykr/backend-test-desent/app"
	"github.com/zikrykr/backend-test-desent/config"
)

// @title Backend Test Desent API
// @version 1.0
// @description This is a swagger for Backend Test Desent
// @BasePath /
func main() {
	fiberApp := app.SetupApp()

	port := config.GetEnv("PORT", "3000")
	appHost := config.GetEnv("APP_HOST", "localhost")

	go func() {
		log.Println("Server is starting on " + appHost + ":" + port)
		if err := fiberApp.Listen(appHost + ":" + port); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Gracefully shutting down server...")

	if err := fiberApp.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server was successful shutdown.")
}
