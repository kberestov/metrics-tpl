package ports

import (
	"github.com/kberestov/metrics-tpl/internal/common/domain"
)

// A MetricStore represents a storage for metrics.
//
//go:generate go run github.com/vektra/mockery/v2@v2.48.0 --name=MetricStore
type MetricStore interface {
	// GetValue returns a value of a metric with the given name.
	// Returns ErrMetricNotFound if no metric found.
	GetValue(n domain.MetricName) (domain.MetricValue, error)

	// SaveValue saves the given value of a metric with the given name.
	// Returns ErrNoMetricValue if the value is not present.
	// Returns ErrMetricValueKindMismatch if the metric already has
	// a value of another kind.
	SaveValue(n domain.MetricName, v domain.MetricValue) error
}
