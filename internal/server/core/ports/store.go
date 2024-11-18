package ports

import (
	"github.com/kberestov/metrics-tpl/internal/common/domain"
)

// A MetricStore represents a storage for metrics.
type MetricStore interface {
	// GetValue returns a value of a metric with the given name.
	// Returns ErrMetricNotFound if no metric found.
	GetValue(n domain.MetricName) (domain.MetricValue, error)

	// SaveValue saves the given value of a metric with the given name.
	SaveValue(n domain.MetricName, v domain.MetricValue) error
}
