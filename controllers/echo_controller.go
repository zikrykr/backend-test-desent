package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/zikrykr/backend-test-desent/services"
	"github.com/zikrykr/backend-test-desent/utils"
)

type EchoController struct {
}

func NewEchoController(healthService services.HealthServiceInterface) EchoControllerInterface {
	return &EchoController{}
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

	if len(reqBody) == 0 {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "request body cannot be empty", nil)
	}

	var resp json.RawMessage
	if err := utils.UnmarshalJSON(reqBody, &resp); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "failed to parse request body to valid JSON", err)
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
