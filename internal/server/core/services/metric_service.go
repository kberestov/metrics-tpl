package services

import (
	"errors"
	"fmt"

	d "github.com/kberestov/metrics-tpl/internal/common/domain"
	p "github.com/kberestov/metrics-tpl/internal/server/core/ports"
)

type MetricService struct {
	store p.MetricStore
}

func NewMetricService(store p.MetricStore) *MetricService {
	return &MetricService{store: store}
}

// TODO: Think about idempotency for the use case.
func (s *MetricService) Update(k string, n string, v string) error {
	const op = "services.MetricService.Update"

	kind, err := d.ParseMetricKind(k)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	name, err := d.ParseMetricName(n)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	switch kind {
	case d.KindCounter:
		// TODO: Here should be some sort of sync for the case
		//       when the same counter is being updated in multiple routines
		//       to avoid data racing.
		updatingVal, err := d.ParseCounterValue(v)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		currentVal, err := s.store.GetCounter(name)
		if err != nil && !errors.Is(err, p.ErrMetricNotFound) {
			return fmt.Errorf("%s: %w", op, err)
		}
		newVal := currentVal + updatingVal
		if err := s.store.SaveCounter(name, newVal); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	case d.KindGauge:
		newVal, err := d.ParseGaugeValue(v)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		if err := s.store.SaveGauge(name, newVal); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}
