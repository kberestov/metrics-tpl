package main

import (
	"fmt"
	"log"

	"github.com/kberestov/metrics-tpl/internal/adapters/config"
	"github.com/kberestov/metrics-tpl/internal/adapters/handlers/api/rest"
	"github.com/kberestov/metrics-tpl/internal/adapters/store"
	"github.com/kberestov/metrics-tpl/internal/core/services"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.Config{ServerAPIAddress: `:8080`}

	memStore := store.NewMemStore()
	metricSvc := services.NewMetricService(memStore)

	api := rest.NewServerAPI(cfg, metricSvc)
	if err := api.Run(); err != nil {
		return fmt.Errorf("failed to run REST API: %w", err)
	}

	return nil
}
