package main

import (
	"fmt"
	"log"

	"github.com/kberestov/metrics-tpl/internal/server/adapters/api/rest"
	"github.com/kberestov/metrics-tpl/internal/server/adapters/stores"
	"github.com/kberestov/metrics-tpl/internal/server/config"
	"github.com/kberestov/metrics-tpl/internal/server/core/services"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.Config{RESTAddress: `:8080`}

	updater := services.NewMetricUpdater(
		stores.NewMemStore(),
	)

	restAPI := rest.New(cfg, updater)

	// TODO: Implement graceful shutdown with reaction on OS signals.
	if err := restAPI.Run(); err != nil {
		return fmt.Errorf("failed to run REST API: %w", err)
	}

	return nil
}
