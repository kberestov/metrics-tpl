package rest

import (
	"fmt"
	"net/http"

	"github.com/kberestov/metrics-tpl/internal/server/config"
	"github.com/kberestov/metrics-tpl/internal/server/core/ports"
)

type pathValue string

const (
	pvMetricKind  pathValue = "kind"
	pvMetricName  pathValue = "name"
	pvMetricValue pathValue = "value"
)

type API struct {
	srv *http.Server
}

func New(cfg config.Config, u ports.MetricUpdater) *API {
	h := &handler{updater: u}

	router := http.NewServeMux()

	router.HandleFunc(
		fmt.Sprintf(
			"POST /update/{%s}/{%s}/{%s}",
			pvMetricKind,
			pvMetricName,
			pvMetricValue,
		),
		h.UpdateMetric,
	)

	return &API{
		srv: &http.Server{
			Addr:    cfg.RESTAddress,
			Handler: router,
		},
	}
}

func (api *API) Run() error {
	if err := api.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to run HTTP server: %w", err)
	}

	return nil
}
