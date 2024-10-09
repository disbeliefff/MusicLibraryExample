package server

import (
	"MusicLibrary/api/routers"
	"MusicLibrary/cfg"
	"MusicLibrary/storage"
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
)

type Server struct {
	srv *echo.Echo
	cfg cfg.Config
	log *slog.Logger
}

func NewServer(cfg cfg.Config, log *slog.Logger) *Server {
	s := &Server{
		srv: echo.New(),
		cfg: cfg,
		log: log,
	}

	routers.SetupRouter(s.srv, &storage.SongStorage{})
	return s
}

func (s *Server) Start() error {
	port := s.cfg.Port
	if port == "" {
		port = "8080"
	}

	s.log.Info("Staring server on port " + port)
	return s.srv.Start(fmt.Sprintf(":%s", port))
}

//func (s *Server) Stop() error {
//	return s.srv.Shutdown()
//}
