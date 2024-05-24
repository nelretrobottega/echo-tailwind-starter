package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nelretrobottega/le-donne/internal"
	"github.com/nelretrobottega/le-donne/internal/product/db"
)

type repository struct {
	q    *db.Queries
	conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) *repository {
	return &repository{
		q:    db.New(conn),
		conn: conn,
	}
}

func (r *repository) GetProducts(ctx context.Context, id int64, limit int32) ([]internal.Product, error) {
	dbProducts, err := r.q.GetProducts(ctx, db.GetProductsParams{
		ID:    id,
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}

	products := make([]internal.Product, len(dbProducts))

	for i, p := range dbProducts {
		products[i] = internal.Product{
			Title:       p.Title,
			Description: p.Description,
			ImgURL:      p.Img,
		}
	}

	return products, nil
}
