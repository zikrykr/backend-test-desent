package services

import (
	"github.com/zikrykr/backend-test-desent/model"
	"github.com/zikrykr/backend-test-desent/utils"
)

type HealthServiceInterface interface {
	CheckHealth() model.Health
}

type BookServiceInterface interface {
	Create(book *model.CreateBookRequest) (*model.Book, *utils.ErrorObj)
	FindByID(id string) (*model.Book, *utils.ErrorObj)
	FindAll() ([]*model.Book, *utils.ErrorObj)
	Update(id string, book *model.UpdateBookRequest) (*model.Book, *utils.ErrorObj)
	Delete(id string) *utils.ErrorObj
}
