package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	//Instancia de echo
	e := echo.New()

	//Route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
