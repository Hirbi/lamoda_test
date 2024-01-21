package product

import "app/internal/models"

type Repository interface {
	ReserveProducts(productCodes map[string]int) error
	ReleaseProducts(productCodes map[string]int) error
	AvailableProducts(storageId string) ([]models.Product, error)
}
