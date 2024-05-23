package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Product struct {
	Title       string
	Description string
}

func IndexHandler(c echo.Context) error {
	data := make(map[string]any)

	data["Name"] = "World"
	data["Products"] = []Product{
		{Title: "Shoe 1", Description: "Very good!"},
		{Title: "Shoe 2", Description: "Very good too!"},
	}

	return c.Render(http.StatusOK, "index.html", data)
}
