package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/nelretrobottega/echo-tailwind-starter/internal/product"
	"github.com/nelretrobottega/echo-tailwind-starter/registry"
	"golang.org/x/net/http2"
)

//go:generate pnpm build

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	e := echo.New()
	e.Renderer = registry.NewTemplateRegistry("views")

	pool, dbConnCleanup, err := newPostgresPool()
	if err != nil {
		logger.Error("failed to connect to db", slog.String("err", err.Error()))
		os.Exit(1)
	}

	defer dbConnCleanup()

	e.Static("/static", "static")

	{
		container := product.Container(pool)
		e.GET("/", container.Index())
	}

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	if err := e.StartH2CServer(":8000", s); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func newPostgresPool() (*pgxpool.Pool, func() error, error) {
	dsn := os.Getenv("POSTGRES_DSN")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)

	if err != nil {
		return nil, nil, err
	}

	cleanup := func() error {
		cancel()
		pool.Close()
		return nil
	}

	return pool, cleanup, nil
}
