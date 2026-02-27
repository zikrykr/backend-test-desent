package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zikrykr/backend-test-desent/config"
	docs "github.com/zikrykr/backend-test-desent/docs"
	"github.com/zikrykr/backend-test-desent/infrastructure"
	"github.com/zikrykr/backend-test-desent/middleware"
	"github.com/zikrykr/backend-test-desent/routes"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @BasePath /
func main() {
	config.LoadEnv()

	port := config.GetEnv("PORT", "3000")
	appHost := config.GetEnv("APP_HOST", "localhost")

	docs.SwaggerInfo.Host = appHost + ":" + port

	infra := infrastructure.Infrastructure{
		Cache: infrastructure.NewCache(),
	}

	app := fiber.New()

	app.Use(middleware.SetupCors())

	routes.SetupRoutes(routes.RouteConfig{
		App:            app,
		Infrastructure: infra,
	})

	log.Println("Server is starting on " + appHost + ":" + port)
	log.Fatal(app.Listen(appHost + ":" + port))
}
