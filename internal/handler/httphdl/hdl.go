package httphdl

import (
	"goredis/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type httphdl struct {
	srv ports.CatalogService
}

func NewCatalogHandler(srv ports.CatalogService) httphdl {
	return httphdl{srv: srv}
}

func (r httphdl) GetProducts(c echo.Context) error {
	products, err := r.srv.GetProducts()
	if err != nil {
		return err
	}
	return c.JSON(200, products)
}
