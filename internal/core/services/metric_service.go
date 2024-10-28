package services

import (
	"fmt"

	"github.com/kberestov/metrics-tpl/internal/core/domain"
	"github.com/kberestov/metrics-tpl/internal/core/ports"
)

type MetricService struct {
	store ports.MetricStore
}

func NewMetricService(store ports.MetricStore) *MetricService {
	return &MetricService{store: store}
}

func (s *MetricService) Update(id *domain.MetricID, newValue domain.MetricValue) (domain.MetricValue, error) {
	// TODO: add smart mutex to sync operations with the same metric IDs

	currentValue, err := s.store.Get(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read value from store: %w", err)
	}

	updatedValue, err := domain.CalcNewMetricValue(id.Kind(), currentValue, newValue)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate updated value: %w", err)
	}

	if err := s.store.Save(id, updatedValue); err != nil {
		return nil, fmt.Errorf("failed to save value to store: %w`", err)
	}

	return updatedValue, nil
}
