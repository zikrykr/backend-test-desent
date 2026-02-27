package services

import (
	"github.com/zikrykr/backend-test-desent/model"
)

type HealthServiceInterface interface {
	CheckHealth() model.Health
}

type BookServiceInterface interface {
	Create(book *model.CreateBookRequest) (*model.Book, error)
}
