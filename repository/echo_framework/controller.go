package echo_framework

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
)

type test struct {
	Somethings string
}

func hello(c echo.Context) error {
	//get Request form value
	v, _ := c.FormParams()
	fmt.Println(v)

	//bind Request to struct
	thing := test{Somethings: "things"}
	err := c.Bind(thing)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.String(http.StatusOK, "Hello, World!")
}
