package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/say-hello/:data-type", sayHelloHandler)
	e.POST("/say-hello", sayHelloPostWithBodyHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

type person struct {
	Name    string "json:'name'"
	Surname string "json:'surname'"
}

func sayHelloHandler(c echo.Context) error {
	dataType := c.Param("data-type")

	name := c.QueryParam("name")
	surname := c.QueryParam("surname")

	if dataType == "string" {
		//http://localhost:8080/say-hello/string?name=Tuna&surname=Aydinoglu
		return c.String(http.StatusOK, fmt.Sprintf("Hello %s %s", name, surname))
	}

	if dataType == "json" {
		//http://localhost:8080/say-hello/json?name=Tuna&surname=Aydinoglu
		return c.JSON(http.StatusOK, map[string]string{
			"name":    name,
			"surname": surname,
		})
	}

	//http://localhost:8080/say-hello/sdf
	return c.String(http.StatusBadRequest, "data-type can be entered as string or json...")
}

func sayHelloPostWithBodyHandler(c echo.Context) error {
	user := person{}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	/*	curl--
		location--
		request
		POST
		'http://localhost:8080/say-hello' \
		--header
		'Content-Type: application/json' \
		--data - raw
		'{
		"name":"Tunahan",
			"surname":"Aydinoglu"
		}'*/

	return c.String(http.StatusOK, fmt.Sprintf("Hello %s %s", user.Name, user.Surname))
}
