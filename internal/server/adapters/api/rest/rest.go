package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kberestov/metrics-tpl/internal/server/config"
	"github.com/kberestov/metrics-tpl/internal/server/core/domain"
	"github.com/kberestov/metrics-tpl/internal/server/core/ports"
)

type pathValue string

const (
	pvMetricKind  pathValue = "kind"
	pvMetricName  pathValue = "name"
	pvMetricValue pathValue = "value"
)

const (
	defaultReadTimeout     = 3 * time.Second
	defaultWriteTimeout    = 3 * time.Second
	defaultShutdownTimeout = 3 * time.Second
)

type API struct {
	srv    *http.Server
	notify chan error
}

func NewAPI(cfg config.Config, ms ports.MetricService) *API {
	h := &handler{metricSvc: ms}

	router := http.NewServeMux()
	router.HandleFunc(
		fmt.Sprintf("POST /update/{%s}/{%s}/{%s}", pvMetricKind, pvMetricName, pvMetricValue),
		h.UpdateMetric,
	)

	return &API{
		srv: &http.Server{
			Addr:         cfg.RESTAddress,
			Handler:      router,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
		},
		notify: make(chan error),
	}
}

func (api *API) Notify() <-chan error {
	return api.notify
}

func (api *API) Run() {
	go func() {
		api.notify <- api.srv.ListenAndServe()
		close(api.notify)
	}()
}

func (api *API) Shutdown() error {
	const op = "adapters.api.rest.Shutdown"

	ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
	defer cancel()

	if err := api.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

type handler struct {
	metricSvc ports.MetricService
}

func (h *handler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	const op = "adapters.api.rest.UpdateMetric"

	logError := func(err error) {
		log.Println(fmt.Errorf("%s: %w", op, err))
	}

	id, err := domain.ParseMetricID(
		r.PathValue(string(pvMetricKind)),
		r.PathValue(string(pvMetricName)),
	)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updating, err := domain.ParseMetricValue(
		id.Kind(),
		r.PathValue(string(pvMetricValue)),
	)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updated, err := h.metricSvc.UpdateValue(id, updating)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(updated.String())); err != nil {
		logError(err)
	}
}
