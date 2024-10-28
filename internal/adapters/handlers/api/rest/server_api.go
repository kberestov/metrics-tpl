package rest

import (
	"fmt"
	"log"
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
	kind := r.PathValue(string(pvKind))
	name := r.PathValue(string(pvName))
	value := r.PathValue(string(pvValue))

	metricID, updateValue, err := domain.ParseMetric(kind, name, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedValue, err := h.metricSvc.Update(metricID, updateValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(updatedValue.String())); err != nil {
		log.Printf("write failed: %v", err)
	}
}
