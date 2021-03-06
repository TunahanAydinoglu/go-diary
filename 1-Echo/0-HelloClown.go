package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", helloWorldHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Clown!")
}
