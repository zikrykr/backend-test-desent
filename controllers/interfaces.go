package controllers

import "github.com/gofiber/fiber/v2"

type HealthControllerInterface interface {
	Check(ctx *fiber.Ctx) error
	Ping(ctx *fiber.Ctx) error
}
