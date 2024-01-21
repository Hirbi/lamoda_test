package product

import "app/internal/models"

type UseCase interface {
	ReserveProducts(productCodes []string) error
	ReleaseProducts(productCodes []string) error
	AvailableProducts(storageId string) ([]models.Product, error)
}
