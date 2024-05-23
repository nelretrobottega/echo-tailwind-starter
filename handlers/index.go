package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Product struct {
	Title       string
	Description string
	ImgURL      string
}

func IndexHandler(c echo.Context) error {
	data := make(map[string]any)

	data["Name"] = "World"
	data["Products"] = []Product{
		{
			Title:       "Shoe 1",
			Description: "Very good!",
			ImgURL:      "/static/img/1.webp",
		},
		{
			Title:       "Shoe 2",
			Description: "Very good too!",
			ImgURL:      "/static/img/2.webp",
		},
	}

	return c.Render(http.StatusOK, "index.html", data)
}
