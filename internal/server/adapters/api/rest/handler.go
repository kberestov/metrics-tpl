package rest

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/kberestov/metrics-tpl/internal/common/domain"
	"github.com/kberestov/metrics-tpl/internal/server/core/ports"
)

type handler struct {
	metricSvc ports.MetricService
}

// POST /update/{pvMetricKind}/{pvMetricName}/{pvMetricValue}
func (h *handler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handler.UpdateMetric"

	logError := func(err error) {
		log.Println(fmt.Errorf("%s: %w", op, err))
	}

	kind := r.PathValue(string(pvMetricKind))
	name := r.PathValue(string(pvMetricName))
	value := r.PathValue(string(pvMetricValue))

	if err := h.metricSvc.Update(kind, name, value); err != nil {
		switch {
		case errors.Is(err, domain.ErrUnknownMetricKind),
			errors.Is(err, domain.ErrInvalidMetricName),
			errors.Is(err, domain.ErrInvalidMetricValue):
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		logError(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
