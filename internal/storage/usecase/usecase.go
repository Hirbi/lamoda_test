package usecase

import (
	"app/internal/storage"
)

type StorageBL struct {
	storageRepo storage.Repository
}

func NewProductUseCase(productRepo storage.Repository) *StorageBL {
	return &StorageBL{
		storageRepo: productRepo,
	}
}
