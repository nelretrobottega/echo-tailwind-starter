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
			ImgURL:      "https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.jpg",
		},
		{
			Title:       "Shoe 2",
			Description: "Very good too!",
			ImgURL:      "https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.jpg",
		},
	}

	return c.Render(http.StatusOK, "index.html", data)
}
