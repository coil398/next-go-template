package main

import (
	"fmt"

	"backend/handler"
	"backend/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func startServer(c *Config, s *store.Store) error {
	e := echo.New()
	registerMiddlewares(e)

	server := &handler.Server{
		Store: s,
	}
	handler.RegisterHandlers(e, server)

	return e.Start(fmt.Sprintf(":%d", c.Server.Port))
}

func registerMiddlewares(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())
}
