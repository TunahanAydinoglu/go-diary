package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(commonHelloMiddleware)
	adminGroup := e.Group("/admin", adminHelloMiddleware)

	e.GET("/", mainIndexHandler)
	adminGroup.GET("", adminMainIndexHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func adminHelloMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("Hello Admin Middleware")
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

func commonHelloMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("Hello Common Middleware")
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

func adminMainIndexHandler(c echo.Context) error {

	return c.String(http.StatusOK, "Welcome admin index page!")
}

func mainIndexHandler(c echo.Context) error {

	return c.String(http.StatusOK, "Welcome index page!")
}
