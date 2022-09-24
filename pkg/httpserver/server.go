package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	defReadTimeout     time.Duration = 5 * time.Second
	defWriteTimeout                  = 5 * time.Second
	defAddr                          = ":80"
	defShutdownTimeout               = 3 * time.Second
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  defReadTimeout,
		WriteTimeout: defWriteTimeout,
		Addr:         defAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error),
		shutdownTimeout: defShutdownTimeout,
	}

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
