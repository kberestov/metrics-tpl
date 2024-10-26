package rest

import (
	"fmt"
	"net/http"

	"github.com/kberestov/metrics-tpl/internal/adapters/config"
	"github.com/kberestov/metrics-tpl/internal/core/ports"
)

type ServerREST struct {
	srv *http.Server
}

func NewServerREST(cfg config.ServerConfig, ms ports.Server) *ServerREST {
	mux := http.NewServeMux()

	return &ServerREST{
		srv: &http.Server{
			Addr:    cfg.RESTAddr,
			Handler: mux,
		},
	}
}

func (api *ServerREST) Run() error {
	if err := api.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to run HTTP server: %w", err)
	}

	return nil
}
