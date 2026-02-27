package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikrykr/backend-test-desent/infrastructure"
	"github.com/zikrykr/backend-test-desent/model"
	"github.com/zikrykr/backend-test-desent/services"
	"github.com/zikrykr/backend-test-desent/utils"
)

type BookController struct {
	BookService services.BookServiceInterface
	Logger      *infrastructure.Logger
}

func NewBookController(
	bookService services.BookServiceInterface,
	logger *infrastructure.Logger,
) BookControllerInterface {
	return &BookController{
		BookService: bookService,
		Logger:      logger,
	}
}

// Create returns a simple create response.
// @Summary Create
// @Description create a new book.
// @Tags book
// @Accept json
// @Produce json
// @Param body body model.CreateBookRequest true "Request body"
// @Success 201 {object} model.Book
// @Router /api/books [post]
func (c *BookController) Create(ctx *fiber.Ctx) error {
	var req model.CreateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.ErrorResponse(ctx, c.Logger, fiber.StatusBadRequest, "failed to parse request body to valid JSON", err)
	}

	bookResult, err := c.BookService.Create(&req)
	if err != nil {
		return utils.ErrorResponse(ctx, c.Logger, fiber.StatusInternalServerError, "failed to create book", err)
	}

	return utils.SuccessPlainResponse(ctx, c.Logger, fiber.StatusCreated, bookResult)
}
