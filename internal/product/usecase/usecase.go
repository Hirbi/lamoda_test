package usecase

import (
	"app/internal/models"
	"app/internal/product"
)

type ProductBL struct {
	productRepo product.Repository
}

func NewProductUseCase(productRepo product.Repository) *ProductBL {
	return &ProductBL{
		productRepo: productRepo,
	}
}

func (p *ProductBL) countCodes(productCodes []string) map[string]int {
	codesCount := make(map[string]int)
	for _, code := range productCodes {
		codesCount[code] = codesCount[code] + 1
	}
	return codesCount
}

func (p *ProductBL) ReserveProducts(productCodes []string) error {
	return p.productRepo.ReserveProducts(p.countCodes(productCodes))
}

func (p *ProductBL) ReleaseProducts(productCodes []string) error {

	return p.productRepo.ReleaseProducts(p.countCodes(productCodes))
}

func (p *ProductBL) AvailableProducts(storageId string) ([]models.Product, error) {
	products, err := p.productRepo.AvailableProducts(storageId)
	return products, err
}
