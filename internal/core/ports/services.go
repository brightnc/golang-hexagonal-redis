package ports

import "goredis/internal/core/domains"

type CatalogService interface {
	GetProducts() ([]domains.ProductSrv, error)
}
