package handler

import (
	"backend/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Store *store.Store
}

func RegisterHandlers(e *echo.Echo, s *Server) {
	e.GET("/healthz", s.healthz)

	v1 := e.Group("v1")
	registerV1UserHandlers(v1, s)
}

func (s *Server) healthz(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
