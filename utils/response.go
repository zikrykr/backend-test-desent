package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikrykr/backend-test-desent/infrastructure"
	"github.com/zikrykr/backend-test-desent/model"
)

func SuccessResponse(ctx *fiber.Ctx, logger *infrastructure.Logger, httpStatus int, message string, data any) error {
	resp := model.Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	logger.Info("Success response: " + message)
	return ctx.Status(httpStatus).JSON(resp)
}

func SuccessPlainResponse(ctx *fiber.Ctx, logger *infrastructure.Logger, httpStatus int, data any) error {
	logger.Info("Success plain response")
	ctx.Set("Content-Type", "application/json")

	if data == nil {
		return ctx.Status(httpStatus).JSON(fiber.Map{})
	}

	return ctx.Status(httpStatus).JSON(data)
}

func ErrorResponse(ctx *fiber.Ctx, logger *infrastructure.Logger, errObj *ErrorObj) error {
	resp := model.Response{
		Success: false,
		Message: errObj.Message,
		Data:    errObj,
	}

	logger.Error("Error response: "+errObj.Message, errObj.Error)
	return ctx.Status(errObj.HTTPStatus).JSON(resp)
}
