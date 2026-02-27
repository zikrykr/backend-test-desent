package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikrykr/backend-test-desent/services"
)

type HealthController struct {
	healthService services.HealthServiceInterface
}

func NewHealthController(healthService services.HealthServiceInterface) HealthControllerInterface {
	return &HealthController{
		healthService: healthService,
	}
}

// Check returns the health status of the API.
// @Summary Check Health
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} model.Health
// @Router /api/health [get]
func (c *HealthController) Check(ctx *fiber.Ctx) error {
	result := c.healthService.CheckHealth()
	return ctx.Status(fiber.StatusOK).JSON(result)
}
