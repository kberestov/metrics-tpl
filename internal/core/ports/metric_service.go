package ports

import "github.com/kberestov/metrics-tpl/internal/core/domain"

type MetricService interface {
	Update(id *domain.MetricID, update domain.MetricValue) (domain.MetricValue, error)
}
