package todo

import (
	"context"
	"net/http"
	"todo-app/pkg/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.HTTPServer, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           cfg.Address,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		Handler:        handler,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		IdleTimeout:    cfg.IdleTimeout,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
