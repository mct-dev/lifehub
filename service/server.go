package service

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// Server an Echo web server
type Server struct {
	// Echo Server
	*echo.Echo
}

// Init create an Echo server with middleware, ui, routes
func Init() *Server {
	server := &Server{}
	server.initEcho()
	server.initMiddleware()

	e.GET("/", home)

	return server
}

func (s *Server) initEcho() {
	s.Echo = echo.New()
	s.HideBanner = true

	// s.HTTPErrorHandler = handlers.HTTPErrorHandler
}

func (s *Server) initMiddleware() {
	s.Echo.Use(echoMiddleware.Logger())
}

// Start start the Echo server
func (s *Server) Start() {
	s.Echo.Logger.Fatal(s.Echo.Start(":1323"))
}
