package ports

import "github.com/kberestov/metrics-tpl/internal/core/domain"

type MetricService interface {
	Update(m domain.Metric) error
}
