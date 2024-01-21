package http

import (
	"app/internal/product"
	"log"
	"net/http"
)

type ProductService struct {
	productBL product.UseCase
}

func NewHandler(useCase product.UseCase) *ProductService {
	return &ProductService{
		productBL: useCase,
	}
}

type ReserveReleaseInput struct {
	Codes []string
}

func (s *ProductService) ReserveProducts(r *http.Request, args *ReserveReleaseInput, reply *interface{}) error {
	log.Printf("Request from %v", r.RemoteAddr)
	if args == nil {
		return product.ErrEmptyData
	}
	if len(args.Codes) == 0 {
		return product.ErrNoCodes
	}
	err := s.productBL.ReserveProducts(args.Codes)
	if err != nil {
		return err
	}
	*reply = "reserved"
	return nil
}

func (s *ProductService) ReleaseProducts(r *http.Request, args *ReserveReleaseInput, reply *interface{}) error {
	log.Printf("Request from %v", r.RemoteAddr)
	if args == nil {
		return product.ErrEmptyData
	}
	if len(args.Codes) == 0 {
		return product.ErrNoCodes
	}
	err := s.productBL.ReleaseProducts(args.Codes)
	if err != nil {
		return err
	}
	*reply = "released"
	return nil
}

type AvailableProductsInput struct {
	StorageId string
}

func (s *ProductService) AvailableProducts(r *http.Request, args *AvailableProductsInput, reply *interface{}) error {
	log.Printf("Request from %v", r.RemoteAddr)
	if args == nil || args.StorageId == "" {
		return product.ErrStorageIdIsNull
	}
	products, err := s.productBL.AvailableProducts(args.StorageId)
	*reply = products
	return err
}
