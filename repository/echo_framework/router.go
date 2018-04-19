package echo_framework

import "github.com/labstack/echo"

func router(e *echo.Echo) {

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":8088"))
}
