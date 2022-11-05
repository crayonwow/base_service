package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type Registrator interface {
	Register(r *mux.Router)
}

type httpServer struct {
	srv *http.Server
	cfg Config
}

func (h *httpServer) Run(ctx context.Context) error {
	l, err := net.Listen("tcp", h.srv.Addr)
	if err != nil {
		return fmt.Errorf("net listen: %w", err)
	}
	h.srv.BaseContext = func(l net.Listener) context.Context {
		return ctx
	}
	go func() {
		sErr := h.srv.Serve(l)
		if errors.Is(sErr, http.ErrServerClosed) {
			panic(err)
		}
	}()

	return nil
}

func (h *httpServer) Stop(_ context.Context) error {
	return h.srv.Close()
}

func register(in registratorIn) http.Handler {
	r := mux.NewRouter()
	for _, registrator := range in.Registrators {
		registrator.Register(r)
	}
	return r
}

func newHTTPServer(h http.Handler, cfg Config) *httpServer {
	return &httpServer{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Handler: h,
		},
		cfg: cfg,
	}
}
