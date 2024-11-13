package ports

import (
	"errors"

	d "github.com/kberestov/metrics-tpl/internal/common/domain"
)

var ErrMetricNotFound = errors.New("metric not found")

// A MetricStore represents a storage for metrics.
type MetricStore interface {
	// GetCounter returns a value of a counter by the given name.
	// Returns ErrMetricNotFound if no metric found.
	GetCounter(name d.MetricName) (d.CounterValue, error)

	// SaveCounter saves the given value to a counter by the given name.
	// If no counter found, the new one must be created.
	SaveCounter(name d.MetricName, value d.CounterValue) error

	// SaveGauge saves the given value to a gauge by the given name.
	// If no gauge found, the new one must be created.
	SaveGauge(name d.MetricName, value d.GaugeValue) error
}
