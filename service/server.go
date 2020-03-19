package service

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/mct-dev/lifehub/config"
)

// Server an Echo web server
type Server struct {
	// Echo Server
	*echo.Echo

	// Config
	config *config.Config

	// Middleware

	// Route Groups
	api *echo.Group
}

// Init create an Echo server with middleware, ui, routes
func Init(config *config.Config) *Server {
	server := &Server{
		config: config,
	}
	server.initEcho()
	server.initMiddleware()
	server.initUI()
	server.initApis()

	return server
}

// Start start the Echo server
func (s *Server) Start() {
	s.Logger.Fatal(s.Echo.Start(":1323"))
}

func (s *Server) initEcho() {
	s.Echo = echo.New()
	s.HideBanner = true

	// s.HTTPErrorHandler = handlers.HTTPErrorHandler
}

func (s *Server) initMiddleware() {
	s.Use(echoMiddleware.Logger())
}

func (s *Server) initUI() {
	// loadUI := s.config.Env == "production"
	// defer logStatus("UI", loadUI)

	// if !loadUI {
	// 	return
	// }

	riceConfig := rice.Config{
		LocateOrder: []rice.LocateMethod{rice.LocateEmbedded},
	}

	uiAssets, err := riceConfig.FindBox("../ui/dist")
	if err != nil {
		panic("Static ui/dist folder not found. Build it with `cd ui && yarn build`.")
	}

	assetHandler := http.FileServer(uiAssets.HTTPBox())
	s.GET("/", echo.WrapHandler(assetHandler))
	s.GET("/favicon*", echo.WrapHandler(assetHandler))
	s.GET("/css/*", echo.WrapHandler(http.StripPrefix("/", assetHandler)))
	s.GET("/js/*", echo.WrapHandler(http.StripPrefix("/", assetHandler)))
	s.GET("/fonts/*", echo.WrapHandler(http.StripPrefix("/", assetHandler)))
	s.GET("/img/*", echo.WrapHandler(http.StripPrefix("/", assetHandler)))

}
