package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/zikrykr/backend-test-desent/services"
	"github.com/zikrykr/backend-test-desent/utils"
)

type EchoController struct {
	healthService services.HealthServiceInterface
}

func NewEchoController(healthService services.HealthServiceInterface) EchoControllerInterface {
	return &EchoController{
		healthService: healthService,
	}
}

// Echo returns a simple echo response.
// @Summary Echo
// @Description echo the request.
// @Tags echo
// @Accept */*
// @Produce json
// @Param body body any true "Request body"
// @Success 200 {object} any
// @Router /api/echo [post]
func (c *EchoController) Echo(ctx *fiber.Ctx) error {
	reqBody := ctx.Body()

	jsonB, err := utils.ParseJSON(reqBody)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "failed to parse request body", err)
	}

	logrus.Debug(string(jsonB))

	var data map[string]any
	if err := utils.UnmarshalJSON(reqBody, &data); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "failed to parse request body", err)
	}

	return ctx.Status(fiber.StatusOK).JSON(data)
}
