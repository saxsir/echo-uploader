package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"io"
	"os"

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

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(fmt.Sprintf("files/%s", file.Filename))
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", file.Filename))
	})

	e.GET("/files", func(c echo.Context) error {
		dir := "files"
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return err
		}

		var paths []string
		for _, file := range files {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}

		return c.HTML(http.StatusOK, fmt.Sprintf("<p>%s.</p>", paths))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
