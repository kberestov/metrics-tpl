package services

import (
	"github.com/kberestov/metrics-tpl/internal/core/domain"
	"github.com/kberestov/metrics-tpl/internal/core/ports"
)

type MetricService struct {
	store ports.MetricStore
}

func NewMetricService(store ports.MetricStore) *MetricService {
	return &MetricService{store: store}
}

func (s *MetricService) UpdateValue(v domain.MetricValue) error {
	return nil
}
