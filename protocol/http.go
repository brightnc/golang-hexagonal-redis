package protocol

import "github.com/labstack/echo/v4"

func ServeREST() {
	e := echo.New()

	e.Start(":8000")
}
