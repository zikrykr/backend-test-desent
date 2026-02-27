package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikrykr/backend-test-desent/config"
	docs "github.com/zikrykr/backend-test-desent/docs"
	"github.com/zikrykr/backend-test-desent/infrastructure"
	"github.com/zikrykr/backend-test-desent/middleware"
	"github.com/zikrykr/backend-test-desent/routes"
)

func SetupApp() *fiber.App {
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

	return app
}
