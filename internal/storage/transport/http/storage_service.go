package http

import (
	"app/internal/storage"
)

type StorageService struct {
	productService storage.UseCase
}

func NewHandler(useCase storage.UseCase) *StorageService {
	return &StorageService{
		productService: useCase,
	}
}
