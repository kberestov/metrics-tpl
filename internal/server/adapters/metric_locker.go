package adapters

import (
	"sync"

	"github.com/kberestov/metrics-tpl/internal/server/core/domain"
)

// MemMetricLocker provides locking based on mutexes stored in the memory.
type MemMetricLocker struct {
	mtxs  map[domain.MetricID]*sync.Mutex
	outer sync.Mutex
}

func NewMemMetricLocker() *MemMetricLocker {
	return &MemMetricLocker{
		mtxs: make(map[domain.MetricID]*sync.Mutex),
	}
}

func (l *MemMetricLocker) Lock(id domain.MetricID) (unlock func()) {
	l.outer.Lock()

	mtx, ok := l.mtxs[id]
	if !ok {
		mtx = new(sync.Mutex)
		l.mtxs[id] = mtx
	}

	l.outer.Unlock()

	mtx.Lock()

	return func() {
		l.outer.Lock()
		defer l.outer.Unlock()

		delete(l.mtxs, id)
		mtx.Unlock()
	}
}
