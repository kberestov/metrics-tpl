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
		return ports.ErrNoMetricValue
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
		newVal, err = calcNewCounterValue(currentVal, updatingVal)
		if err != nil {
			return err
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

func calcNewCounterValue(curr domain.MetricValue, upd domain.CounterValue) (domain.MetricValue, error) {
	if curr == nil {
		return upd, nil
	}

	c, ok := curr.(domain.CounterValue)
	if !ok {
		return nil, ports.ErrMetricValueKindMismatch
	}

	return c + upd, nil
}
