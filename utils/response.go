package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikrykr/backend-test-desent/model"
)

func SuccessResponse(ctx *fiber.Ctx, httpStatus int, message string, data any) error {
	resp := model.Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	return ctx.Status(httpStatus).JSON(resp)
}
