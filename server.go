package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "static")

	e.POST("/upload", func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", file.Filename))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
