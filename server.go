package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Static("/static", "static")
	e.File("/", "static/index.html")

	e.Logger.Fatal(e.Start(":1323"))
}
