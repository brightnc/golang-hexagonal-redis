package dto

import "goredis/internal/core/domains"

type ProductDto struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func (r ProductDto) ToProductDomain() domains.ProductSrv {
	return domains.ProductSrv{
		ID:       r.ID,
		Name:     r.Name,
		Quantity: r.Quantity,
	}
}
