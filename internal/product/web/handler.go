package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nelretrobottega/le-donne/internal"
)

type handler struct {
	r internal.ProductRepository
}

func NewHandler(r internal.ProductRepository) *handler {
	return &handler{
		r: r,
	}
}

func (h *handler) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")

		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			limitInt = 50
		}

		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			offsetInt = 0
		}

		products, err := h.r.GetProducts(
			c.Request().Context(),
			int64(offsetInt),
			int32(limitInt),
		)
		if err != nil {
			echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		data := make(map[string]any)
		data["Name"] = "World"
		data["Products"] = products

		return c.Render(http.StatusOK, "index.html", data)
	}
}
