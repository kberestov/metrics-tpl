package services

import "github.com/kberestov/metrics-tpl/internal/core/ports"

type Server struct {
	store ports.Store
}

func NewServer(store ports.Store) *Server {
	return &Server{
		store: store,
	}
}
