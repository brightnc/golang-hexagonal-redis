package protocol

import (
	"goredis/internal/handler/httphdl"

	"github.com/labstack/echo/v4"
)

func ServeREST() {
	e := echo.New()
	custHttp := httphdl.NewCatalogHandler(app.svr)
	e.GET("/products", custHttp.GetProducts)

	e.Start(":8000")
}
