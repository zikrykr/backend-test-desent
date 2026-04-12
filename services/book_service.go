package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/zikrykr/backend-test-desent/infrastructure"
	"github.com/zikrykr/backend-test-desent/model"
	"github.com/zikrykr/backend-test-desent/utils"
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

func (s *bookService) Create(req *model.CreateBookRequest) (*model.Book, *utils.ErrorObj) {
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

func (s *bookService) FindByID(id string) (*model.Book, *utils.ErrorObj) {
	cacheKey := fmt.Sprintf("books:%s", id)

	book, found := s.cache.Get(cacheKey)
	if !found {
		return nil, utils.NotFoundError(s.logger, "book not found", nil)
	}

	b, err := utils.ConvertToPtr[model.Book](book)
	if err != nil {
		return nil, utils.InternalServerError(s.logger, "failed to convert book", err)
	}

	s.logger.Info("Book found: " + id)
	return b, nil
}

func (s *bookService) FindAll() ([]*model.Book, *utils.ErrorObj) {
	var bookResult []*model.Book
	books, found := s.cache.GetAll("books:*")
	if !found {
		return nil, utils.NotFoundError(s.logger, "books not found", nil)
	}

	for _, book := range books.([]any) {
		if b, err := utils.ConvertToPtr[model.Book](book); err == nil {
			bookResult = append(bookResult, b)
		}
	}

	s.logger.Info("Books found")
	return bookResult, nil
}

func (s *bookService) Update(id string, req *model.UpdateBookRequest) (*model.Book, *utils.ErrorObj) {
	cacheKey := fmt.Sprintf("books:%s", id)

	book, found := s.cache.Get(cacheKey)
	if !found {
		return nil, utils.NotFoundError(s.logger, "book not found", nil)
	}

	b, err := utils.ConvertToPtr[model.Book](book)
	if err != nil {
		return nil, utils.NotFoundError(s.logger, "book not found", nil)
	}

	b.Title = req.Title
	b.Author = req.Author
	b.Year = req.Year

	s.cache.Set(cacheKey, b, 0)

	s.logger.Info("Book updated: " + id)
	return b, nil
}

func (s *bookService) Delete(id string) *utils.ErrorObj {
	cacheKey := fmt.Sprintf("books:%s", id)

	s.cache.Delete(cacheKey)

	s.logger.Info("Book deleted: " + id)
	return nil
}
