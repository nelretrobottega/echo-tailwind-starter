package product

import (
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nelretrobottega/echo-tailwind-starter/internal"
	"github.com/nelretrobottega/echo-tailwind-starter/internal/product/repository"
	handler "github.com/nelretrobottega/echo-tailwind-starter/internal/product/web"
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
