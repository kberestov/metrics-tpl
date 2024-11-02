package ports

import (
	"github.com/kberestov/metrics-tpl/internal/server/core/domain"
)

// A MetricService represents use cases related to metrics.
type MetricService interface {
	// UpdateValue tries to update a metric by the given ID with the updating value.
	// Returns ErrMetricValueRequired if the updating value is nil.
	UpdateValue(id domain.MetricID, updating domain.MetricValue) (updated domain.MetricValue, err error)
}
