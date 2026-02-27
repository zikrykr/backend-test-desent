package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/zikrykr/backend-test-desent/controllers"
	"github.com/zikrykr/backend-test-desent/infrastructure"
	"github.com/zikrykr/backend-test-desent/services"
)

type RouteConfig struct {
	App            *fiber.App
	Infrastructure infrastructure.Infrastructure
}

func SetupRoutes(config RouteConfig) {
	healthService := services.NewHealthService(config.Infrastructure.Cache)

	healthController := controllers.NewHealthController(healthService, config.Infrastructure.Logger)
	echoController := controllers.NewEchoController(healthService, config.Infrastructure.Logger)

	api := config.App.Group("/api")

	// Swagger Docs
	config.App.Get("/swagger/*", swagger.HandlerDefault)

	api.Get("/health", healthController.Check)
	api.Get("/ping", healthController.Ping)
	api.Post("/echo", echoController.Echo)

	bookService := services.NewBookService(config.Infrastructure.Cache, config.Infrastructure.Logger)
	bookController := controllers.NewBookController(bookService, config.Infrastructure.Logger)
	api.Post("/books", bookController.Create)

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! Welcome to the Fiber backend.")
	})
}
