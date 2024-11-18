package ports

import (
	"github.com/kberestov/metrics-tpl/internal/common/domain"
)

// MetricServerClient represents a client to communicate with the metric server.
type MetricServerClient interface {
	UpdateMetric(n domain.MetricName, v domain.MetricValue) error
}
