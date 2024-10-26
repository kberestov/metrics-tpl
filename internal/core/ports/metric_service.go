package ports

import "github.com/kberestov/metrics-tpl/internal/core/domain"

type MetricService interface {
	UpdateValue(v domain.MetricValue) error
}
