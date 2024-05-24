package internal

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Product struct {
	Title       string
	Description string
	ImgURL      string
}

type ProductRepository interface {
	GetProducts(ctx context.Context, id int64, limit int32) ([]Product, error)
}

type ProductHandler interface {
	Index() echo.HandlerFunc
}
