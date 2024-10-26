package rest

import (
	"fmt"
	"net/http"

	"github.com/kberestov/metrics-tpl/internal/adapters/config"
	"github.com/kberestov/metrics-tpl/internal/core/domain"
	"github.com/kberestov/metrics-tpl/internal/core/ports"
)

type pathValue string

const (
	pvKind  pathValue = "kind"
	pvName  pathValue = "name"
	pvValue pathValue = "value"
)

type ServerAPI struct {
	srv *http.Server
}

func NewServerAPI(cfg config.Config, metricSvc ports.MetricService) *ServerAPI {
	h := &serverHandler{metricSvc: metricSvc}

	mux := http.NewServeMux()
	mux.HandleFunc(
		fmt.Sprintf("POST /update/{%s}/{%s}/{%s}", pvKind, pvName, pvValue),
		h.UpdateMetric)

	return &ServerAPI{
		srv: &http.Server{
			Addr:    cfg.ServerAPIAddress,
			Handler: mux,
		},
	}
}

func (api *ServerAPI) Run() error {
	if err := api.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to run HTTP server: %w", err)
	}
	return nil
}

type serverHandler struct {
	metricSvc ports.MetricService
}

func (h *serverHandler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	k := r.PathValue(string(pvKind))
	n := r.PathValue(string(pvName))
	metric, err := domain.ParseMetric(k, n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := r.PathValue(string(pvValue))
	value, err := metric.ParseValue(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.metricSvc.UpdateValue(value); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
