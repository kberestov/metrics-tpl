package ports

import "github.com/kberestov/metrics-tpl/internal/core/domain"

type MetricStore interface {
	Get(id *domain.MetricID) (domain.MetricValue, error)
	Save(id *domain.MetricID, value domain.MetricValue) error
}
