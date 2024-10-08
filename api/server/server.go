package server

import (
	"MusicLibrary/api/routers"
	"MusicLibrary/cfg"
	"fmt"
	"github.com/labstack/echo/v4"
)

type Server struct {
	srv *echo.Echo
	cfg cfg.Config
}

func NewServer(cfg cfg.Config) *Server {
	s := &Server{
		srv: echo.New(),
		cfg: cfg,
	}

	routers.SetupRouter(s.srv)
	return s
}

func (s *Server) Start() error {
	port := s.cfg.Port
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port %s\n", port)
	return s.srv.Start(fmt.Sprintf(":%s", port))
}
