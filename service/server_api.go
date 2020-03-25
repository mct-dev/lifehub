package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) initApis() {
	s.api = s.Group("/api/v1")
	s.api.POST("/", commandHandler)

	s.registerEventsRoutes()
}

func (s *Server) registerEventsRoutes() {
	s.api.GET("/events", s.getAllEvents)
	s.api.POST("/events", s.addEvent)
}

func commandHandler(c echo.Context) error {
	return c.String(http.StatusOK, "heyyyy")
}

func (s *Server) addEvent(c echo.Context) error {
	// TODO:
	return nil
}

func (s *Server) getAllEvents(c echo.Context) error {
	events, err := s.db.AllEvents()
	if err != nil {
		panic(err)
	}

	for _, ev := range events {
		fmt.Printf("%v", ev)
	}

	return c.JSON(http.StatusOK, events)
}
