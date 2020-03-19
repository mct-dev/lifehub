package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) initApis() {
	s.api = s.Group("/api/v1")
	s.api.POST("/", commandHandler)
}

func commandHandler(c echo.Context) error {
	return c.String(http.StatusOK, "heyyyy")
}
