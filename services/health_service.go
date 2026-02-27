package services

import (
	"github.com/zikrykr/backend-test-desent/infrastructure"
	"github.com/zikrykr/backend-test-desent/model"
)

type healthService struct {
	cache *infrastructure.Cache
}

func NewHealthService(cache *infrastructure.Cache) HealthServiceInterface {
	return &healthService{
		cache: cache,
	}
}

func (s *healthService) CheckHealth() model.Health {
	return model.Health{
		Status:  "success",
		Message: "API is healthy and running",
	}
}
