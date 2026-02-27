package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikrykr/backend-test-desent/infrastructure"
	"github.com/zikrykr/backend-test-desent/services"
	"github.com/zikrykr/backend-test-desent/utils"
)

type HealthController struct {
	Logger        *infrastructure.Logger
	healthService services.HealthServiceInterface
}

func NewHealthController(
	healthService services.HealthServiceInterface,
	logger *infrastructure.Logger,
) HealthControllerInterface {
	return &HealthController{
		healthService: healthService,
		Logger:        logger,
	}
}

// Check returns the health status of the API.
// @Summary Check Health
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/health [get]
func (c *HealthController) Check(ctx *fiber.Ctx) error {
	result := c.healthService.CheckHealth()
	return utils.SuccessResponse(ctx, c.Logger, fiber.StatusOK, "success", result)
}

// Ping returns a simple ping response.
// @Summary Ping
// @Description ping the server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/ping [get]
func (c *HealthController) Ping(ctx *fiber.Ctx) error {
	return utils.SuccessResponse(ctx, c.Logger, fiber.StatusOK, "", nil)
}
