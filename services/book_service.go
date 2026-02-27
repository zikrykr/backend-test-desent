package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/zikrykr/backend-test-desent/infrastructure"
	"github.com/zikrykr/backend-test-desent/model"
)

type bookService struct {
	cache  *infrastructure.Cache
	logger *infrastructure.Logger
}

func NewBookService(cache *infrastructure.Cache, logger *infrastructure.Logger) BookServiceInterface {
	return &bookService{
		cache:  cache,
		logger: logger,
	}
}

func (s *bookService) Create(req *model.CreateBookRequest) (*model.Book, error) {
	bookID := uuid.New().String()

	bookResult := model.Book{
		ID:     bookID,
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}

	cacheKey := fmt.Sprintf("books:%s", bookID)

	s.cache.Set(cacheKey, bookResult, 0)

	s.logger.Info("Book created: " + bookID)
	return &bookResult, nil
}
