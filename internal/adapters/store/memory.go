package store

import (
	"sync"

	"github.com/kberestov/metrics-tpl/internal/core/domain"
)

type MemStore struct {
	values map[domain.MetricID]domain.MetricValue
	mtx    sync.RWMutex
}

func NewMemStore() *MemStore {
	return &MemStore{
		values: make(map[domain.MetricID]domain.MetricValue),
	}
}

func (ms *MemStore) Get(id *domain.MetricID) (domain.MetricValue, error) {
	ms.mtx.RLock()
	defer ms.mtx.RUnlock()

	return ms.values[*id], nil
}

func (ms *MemStore) Save(id *domain.MetricID, value domain.MetricValue) error {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()

	ms.values[*id] = value

	return nil
}
