package main

import (
	"time"

	"github.com/kberestov/metrics-tpl/internal/agent/adapters/clients"
	"github.com/kberestov/metrics-tpl/internal/agent/adapters/scheduler"
	"github.com/kberestov/metrics-tpl/internal/agent/config"
	"github.com/kberestov/metrics-tpl/internal/agent/core/services"
)

func main() {
	cfg := config.Config{
		PollInterval:   2 * time.Second,
		ReportInterval: 10 * time.Second,
		ServerREST: config.ServerREST{
			Host:   "localhost:8080",
			Scheme: "http",
		},
	}

	poller := services.NewMetricPoller()

	reporter := services.NewMetricReporter(
		clients.NewServerRESTClient(cfg),
	)

	sch := scheduler.New(cfg, poller, reporter)

	// TODO: Implement graceful shutdown and reaction on OS signals.
	sch.Run()
}
