package product

import (
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nelretrobottega/le-donne/internal"
	"github.com/nelretrobottega/le-donne/internal/product/repository"
	handler "github.com/nelretrobottega/le-donne/internal/product/web"
)

var (
	repo internal.ProductRepository
	hand internal.ProductHandler

	repositoryOnce sync.Once
	handlerOnce    sync.Once
)

func provideRepository(conn *pgxpool.Pool) internal.ProductRepository {
	repositoryOnce.Do(func() {
		repo = repository.NewRepository(conn)
	})
	return repo
}

func provideHandler(r internal.ProductRepository) internal.ProductHandler {
	handlerOnce.Do(func() {
		hand = handler.NewHandler(r)
	})
	return hand
}
