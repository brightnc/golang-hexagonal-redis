package ports

import "goredis/internal/core/domains"

type ProductRepository interface {
	GetProduct() ([]domains.Product, error)
}
