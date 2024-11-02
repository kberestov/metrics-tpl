package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kberestov/metrics-tpl/internal/server/adapters"
	"github.com/kberestov/metrics-tpl/internal/server/adapters/api/rest"
	"github.com/kberestov/metrics-tpl/internal/server/adapters/store"
	"github.com/kberestov/metrics-tpl/internal/server/config"
	"github.com/kberestov/metrics-tpl/internal/server/core/services"
)

func main() {
	log.Println("server: starting")

	cfg := config.Config{RESTAddress: `:8080`}

	metricSvc := services.NewMetricService(
		store.NewMemStore(),
		adapters.NewMemMetricLocker(),
	)

	restAPI := rest.NewAPI(cfg, metricSvc)

	log.Printf("rest api: starting on %s\n", cfg.RESTAddress)
	restAPI.Run()

	chIntSig := make(chan os.Signal, 1)
	defer close(chIntSig)
	signal.Notify(chIntSig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-chIntSig:
		log.Printf("server: interrupted by signal: %v\n", sig)
	case err := <-restAPI.Notify():
		if !errors.Is(err, http.ErrServerClosed) {
			log.Println(fmt.Errorf("rest api: error occured: %w", err))
		}
	}

	log.Println("rest api: shutting down")

	err := restAPI.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("rest api: failed to shutdown: %w", err))
	}

	log.Println("server: done")
}
