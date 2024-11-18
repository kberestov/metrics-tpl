package rest

import (
	"log"
	"net/http"

	"github.com/kberestov/metrics-tpl/internal/common/domain"
	"github.com/kberestov/metrics-tpl/internal/server/core/ports"
)

type handler struct {
	updater ports.MetricUpdater
}

// UpdateMetric handles the request:
// POST /update/{pvMetricKind}/{pvMetricName}/{pvMetricValue}
func (h *handler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	k, err := domain.ParseMetricKind(r.PathValue(string(pvMetricKind)))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	n, err := domain.ParseMetricName(r.PathValue(string(pvMetricName)))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v, err := domain.ParseMetricValue(k, r.PathValue(string(pvMetricValue)))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("updating metric \"%v\" with value \"%v\"", n, v)
	if err := h.updater.Update(n, v); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
