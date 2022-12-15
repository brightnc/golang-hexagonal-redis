package services

import (
	"goredis/internal/core/domains"
	"goredis/internal/core/ports"
)

type catalogService struct {
	ProductsRepo ports.ProductRepository
}

func NewCatalogService(productsRepo ports.ProductRepository) ports.CatalogService {
	return catalogService{productsRepo}
}

func (r catalogService) GetProducts() (products []domains.ProductSrv, err error) {
	productsDB, err := r.ProductsRepo.GetProduct()
	if err != nil {
		return nil, err
	}
	for _, p := range productsDB {
		products = append(products, domains.ProductSrv{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}
	return products, nil
}
