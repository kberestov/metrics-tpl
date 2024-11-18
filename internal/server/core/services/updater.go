package services

import (
	"errors"
	"fmt"

	"github.com/kberestov/metrics-tpl/internal/common/domain"
	"github.com/kberestov/metrics-tpl/internal/server/core/ports"
)

type MetricUpdater struct {
	store ports.MetricStore
}

func NewMetricUpdater(s ports.MetricStore) *MetricUpdater {
	return &MetricUpdater{store: s}
}

// TODO: Think about idempotency for the use case.
func (u *MetricUpdater) Update(n domain.MetricName, v domain.MetricValue) error {
	if v == nil {
		return errors.New("no metric value provided")
	}

	var newVal domain.MetricValue

	switch v.Kind() {
	case domain.KindCounter:
		// TODO: Here should be some sort of sync for the case
		//       when the same counter is being updated in multiple routines
		//       to avoid data racing.
		updatingVal := v.(domain.CounterValue)

		currentVal, err := u.store.GetValue(n)
		if err != nil {
			if !errors.Is(err, ports.ErrMetricNotFound) {
				return fmt.Errorf("failed to get metric value: %w", err)
			}
		}

		if currentVal == nil {
			newVal = updatingVal
		} else {
			newVal = currentVal.(domain.CounterValue) + updatingVal
		}
	case domain.KindGauge:
		newVal = v.(domain.GaugeValue)
	default:
		return domain.ErrUnknownMetricKind
	}

	if err := u.store.SaveValue(n, newVal); err != nil {
		return fmt.Errorf("failed to save metric value: %w", err)
	}

	return nil
}
