package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:generate pnpm build

func main() {
	e := echo.New()
	e.Renderer = NewTemplateRegistry("views")

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "World")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
