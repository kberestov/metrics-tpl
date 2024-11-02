package ports

import "github.com/kberestov/metrics-tpl/internal/server/core/domain"

// A MetricLocker allows to sync mutating operations for the same metric
// to prevent from data racing.
type MetricLocker interface {
	Lock(id domain.MetricID) (unlock func())
}
