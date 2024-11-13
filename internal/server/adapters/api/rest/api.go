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

func NewAPI(cfg config.Config, metricSvc ports.MetricService) *API {
	handler := &handler{metricSvc: metricSvc}

	router := http.NewServeMux()
	router.HandleFunc(
		fmt.Sprintf("POST /update/{%s}/{%s}/{%s}", pvMetricKind, pvMetricName, pvMetricValue),
		handler.UpdateMetric,
	)

	return &API{
		srv: &http.Server{
			Addr:    cfg.RESTAddress,
			Handler: router,
		},
	}
}

func (api *API) Run() error {
	const op = "rest.API.Run"

	if err := api.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
