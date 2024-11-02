package services

import (
	"errors"
	"fmt"

	"github.com/kberestov/metrics-tpl/internal/server/core/domain"
	"github.com/kberestov/metrics-tpl/internal/server/core/ports"
)

type MetricService struct {
	db     ports.MetricStore
	locker ports.MetricLocker
}

func NewMetricService(s ports.MetricStore, l ports.MetricLocker) *MetricService {
	return &MetricService{
		db:     s,
		locker: l,
	}
}

func (srv *MetricService) UpdateValue(id domain.MetricID, updating domain.MetricValue) (
	updated domain.MetricValue,
	err error,
) {
	const op = "core.services.MetricService.UpdateValue"

	if updating == nil {
		return nil, ports.ErrMetricValueRequired
	}

	unlock := srv.locker.Lock(id)
	defer unlock()

	current, err := srv.db.GetValue(id)
	if err != nil {
		if !errors.Is(err, ports.ErrMetricNotFound) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		current = nil
	}

	updated, err = domain.CalcNewMetricValue(id.Kind(), current, updating)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := srv.db.SaveValue(id, updated); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return updated, nil
}
