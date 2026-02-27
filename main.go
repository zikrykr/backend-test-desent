package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zikrykr/backend-test-desent/config"
	docs "github.com/zikrykr/backend-test-desent/docs"
	"github.com/zikrykr/backend-test-desent/infrastructure"
	"github.com/zikrykr/backend-test-desent/middleware"
	"github.com/zikrykr/backend-test-desent/routes"
)

// @title Backend Test Desent API
// @version 1.0
// @description This is a swagger for Backend Test Desent
// @BasePath /
func main() {
	config.LoadEnv()

	port := config.GetEnv("PORT", "3000")
	appHost := config.GetEnv("APP_HOST", "localhost")

	docs.SwaggerInfo.Host = appHost + ":" + port

	infra := infrastructure.Infrastructure{
		Cache: infrastructure.NewCache(),
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(middleware.SetupCors())

	routes.SetupRoutes(routes.RouteConfig{
		App:            app,
		Infrastructure: infra,
	})

	go func() {
		log.Println("Server is starting on " + appHost + ":" + port)
		if err := app.Listen(appHost + ":" + port); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Gracefully shutting down server...")

	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server was successful shutdown.")
}
