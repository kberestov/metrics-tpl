package ports

import "github.com/kberestov/metrics-tpl/internal/server/core/domain"

// A MetricStore stores metric values.
type MetricStore interface {
	// GetValue tries to find a value of a metric by the given ID.
	// Returns ErrMetricNotFound if the metric is not found.
	GetValue(id domain.MetricID) (domain.MetricValue, error)

	// SaveValue tries to save the value of a metric by the given ID.
	// Does nothing if the value is nil.
	// If the metric doesn't exist, it will be added to the store.
	SaveValue(id domain.MetricID, value domain.MetricValue) error
}
