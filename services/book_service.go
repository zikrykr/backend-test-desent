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

func (s *bookService) FindByID(id string) (*model.Book, error) {
	cacheKey := fmt.Sprintf("books:%s", id)

	book, found := s.cache.Get(cacheKey)
	if !found {
		return nil, fmt.Errorf("book not found")
	}

	b, err := utils.ConvertToPtr[model.Book](book)
	if err == nil {
		return b, nil
	}

	s.logger.Info("Book found: " + id)
	return nil, fmt.Errorf("book not found")
}

func (s *bookService) FindAll() ([]*model.Book, error) {
	var bookResult []*model.Book
	books, found := s.cache.GetAll("books:*")
	if !found {
		return nil, fmt.Errorf("books not found")
	}

	for _, book := range books.([]any) {
		if b, err := utils.ConvertToPtr[model.Book](book); err == nil {
			bookResult = append(bookResult, b)
		}
	}

	s.logger.Info("Books found")
	return bookResult, nil
}

func (s *bookService) Update(id string, req *model.UpdateBookRequest) (*model.Book, error) {
	cacheKey := fmt.Sprintf("books:%s", id)

	book, found := s.cache.Get(cacheKey)
	if !found {
		return nil, fmt.Errorf("book not found")
	}

	b, err := utils.ConvertToPtr[model.Book](book)
	if err != nil {
		return nil, fmt.Errorf("book not found")
	}

	b.Title = req.Title
	b.Author = req.Author
	b.Year = req.Year

	s.cache.Set(cacheKey, b, 0)

	s.logger.Info("Book updated: " + id)
	return b, nil
}

func (s *bookService) Delete(id string) error {
	cacheKey := fmt.Sprintf("books:%s", id)

	s.cache.Delete(cacheKey)

	s.logger.Info("Book deleted: " + id)
	return nil
}
