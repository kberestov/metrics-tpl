package store

import (
	"sync"

	"github.com/kberestov/metrics-tpl/internal/server/core/domain"
	"github.com/kberestov/metrics-tpl/internal/server/core/ports"
)

type MemStore struct {
	values map[domain.MetricKind]map[domain.MetricName]domain.MetricValue
	mtx    sync.RWMutex
}

func NewMemStore() *MemStore {
	return &MemStore{
		values: make(map[domain.MetricKind]map[domain.MetricName]domain.MetricValue),
	}
}

func (s *MemStore) GetValue(id domain.MetricID) (domain.MetricValue, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	byKind, ok := s.values[id.Kind()]
	if !ok {
		return nil, ports.ErrMetricNotFound
	}

	m, ok := byKind[id.Name()]
	if !ok {
		return nil, ports.ErrMetricNotFound
	}

	return m, nil
}

func (s *MemStore) SaveValue(id domain.MetricID, value domain.MetricValue) error {
	if value == nil {
		return nil
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	byKind, ok := s.values[id.Kind()]
	if !ok {
		s.values[id.Kind()] = map[domain.MetricName]domain.MetricValue{
			id.Name(): value,
		}
		return nil
	}

	byKind[id.Name()] = value
	return nil
}
