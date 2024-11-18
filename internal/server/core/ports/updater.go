package ports

import "github.com/kberestov/metrics-tpl/internal/common/domain"

// MetricUpdater represents an application service
// intended to update metrics collected by the server.
type MetricUpdater interface {
	Update(n domain.MetricName, v domain.MetricValue) error
}
