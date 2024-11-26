package ports

import "github.com/kberestov/metrics-tpl/internal/common/domain"

// MetricUpdater represents an application service
// intended to update metrics collected by the server.
//
//go:generate go run github.com/vektra/mockery/v2@v2.48.0 --name=MetricUpdater
type MetricUpdater interface {
	Update(n domain.MetricName, v domain.MetricValue) error
}
