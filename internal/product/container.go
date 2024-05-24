package product

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nelretrobottega/le-donne/internal"
)

func Container(conn *pgxpool.Pool) internal.ProductHandler {
	r := provideRepository(conn)
	h := provideHandler(r)

	return h
}
