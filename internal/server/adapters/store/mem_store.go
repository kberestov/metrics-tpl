package store

import (
	"sync"

	d "github.com/kberestov/metrics-tpl/internal/common/domain"
	p "github.com/kberestov/metrics-tpl/internal/server/core/ports"
)

type MemStore struct {
	values map[d.MetricKind]map[d.MetricName]any
	m      sync.RWMutex
}

func NewMemStore() *MemStore {
	return &MemStore{
		values: make(map[d.MetricKind]map[d.MetricName]any),
	}
}

func (s *MemStore) GetCounter(name d.MetricName) (d.CounterValue, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	counters, ok := s.values[d.KindCounter]
	if !ok {
		var zero d.CounterValue
		return zero, p.ErrMetricNotFound
	}

	value, ok := counters[name]
	if !ok {
		var zero d.CounterValue
		return zero, p.ErrMetricNotFound
	}

	return value.(d.CounterValue), nil
}

func (s *MemStore) SaveCounter(name d.MetricName, value d.CounterValue) error {
	return s.save(d.KindCounter, name, value)
}

func (s *MemStore) SaveGauge(name d.MetricName, value d.GaugeValue) error {
	return s.save(d.KindGauge, name, value)
}

func (s *MemStore) save(kind d.MetricKind, name d.MetricName, value any) error {
	s.m.Lock()
	defer s.m.Unlock()

	byKind, ok := s.values[kind]
	if !ok {
		s.values[kind] = map[d.MetricName]any{
			name: value,
		}
		return nil
	}

	byKind[name] = value
	return nil
}
