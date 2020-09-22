package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/api/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "UP")
	})
	e.Logger.Fatal(e.Start(":1234"))
}
