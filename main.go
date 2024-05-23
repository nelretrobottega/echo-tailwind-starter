package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nelretrobottega/le-donne/handlers"
	"golang.org/x/net/http2"
)

//go:generate pnpm build

func main() {
	e := echo.New()
	e.Renderer = NewTemplateRegistry("views")

	e.Static("/static", "static")

	e.GET("/", handlers.IndexHandler)

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	if err := e.StartH2CServer(":8000", s); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
