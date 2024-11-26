package stores

import (
	"sync"

	"github.com/kberestov/metrics-tpl/internal/common/domain"
	"github.com/kberestov/metrics-tpl/internal/server/core/ports"
)

type MemStore struct {
	vals map[domain.MetricName]domain.MetricValue
	mtx  sync.RWMutex
}

func NewMemStore() *MemStore {
	return &MemStore{
		vals: make(map[domain.MetricName]domain.MetricValue),
	}
}

func (s *MemStore) GetValue(n domain.MetricName) (domain.MetricValue, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	val, ok := s.vals[n]
	if !ok {
		return nil, ports.ErrMetricNotFound
	}

	return val, nil
}

func (s *MemStore) SaveValue(n domain.MetricName, v domain.MetricValue) error {
	if v == nil {
		return ports.ErrNoMetricValue
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	currVal, ok := s.vals[n]
	if ok && currVal.Kind() != v.Kind() {
		return ports.ErrMetricValueKindMismatch
	}

	s.vals[n] = v

	return nil
}
