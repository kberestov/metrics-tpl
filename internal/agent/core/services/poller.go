package services

import (
	"math/rand/v2"
	"runtime"

	"github.com/kberestov/metrics-tpl/internal/agent/core/domain/metrics"
)

type MetricPoller struct {
}

func NewMetricPoller() *MetricPoller {
	return &MetricPoller{}
}

func (p *MetricPoller) Poll() {
	// RandomValue
	metrics.RandomValue.Set(rand.Float64())

	// Runtime metrics:
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)

	// Alloc
	metrics.Alloc.Set(float64(memStats.Alloc))
	metrics.PollCount.Add(1)

	// BuckHashSys
	metrics.BuckHashSys.Set(float64(memStats.BuckHashSys))
	metrics.PollCount.Add(1)

	// Frees
	metrics.Frees.Set(float64(memStats.Frees))
	metrics.PollCount.Add(1)

	// GCCPUFraction
	metrics.GCCPUFraction.Set(memStats.GCCPUFraction)
	metrics.PollCount.Add(1)

	// GCSys
	metrics.GCSys.Set(float64(memStats.GCSys))
	metrics.PollCount.Add(1)

	// HeapAlloc
	metrics.HeapAlloc.Set(float64(memStats.HeapAlloc))
	metrics.PollCount.Add(1)

	// HeapIdle
	metrics.HeapIdle.Set(float64(memStats.HeapIdle))
	metrics.PollCount.Add(1)

	// HeapInuse
	metrics.HeapInuse.Set(float64(memStats.HeapInuse))
	metrics.PollCount.Add(1)

	// HeapObjects
	metrics.HeapObjects.Set(float64(memStats.HeapObjects))
	metrics.PollCount.Add(1)

	// HeapReleased
	metrics.HeapReleased.Set(float64(memStats.HeapReleased))
	metrics.PollCount.Add(1)

	// HeapSys
	metrics.HeapSys.Set(float64(memStats.HeapSys))
	metrics.PollCount.Add(1)

	// LastGC
	metrics.LastGC.Set(float64(memStats.LastGC))
	metrics.PollCount.Add(1)

	// Lookups
	metrics.Lookups.Set(float64(memStats.Lookups))
	metrics.PollCount.Add(1)

	// MCacheInuse
	metrics.MCacheInuse.Set(float64(memStats.MCacheInuse))
	metrics.PollCount.Add(1)

	// MCacheSys
	metrics.MCacheSys.Set(float64(memStats.MCacheSys))
	metrics.PollCount.Add(1)

	// MSpanInuse
	metrics.MSpanInuse.Set(float64(memStats.MSpanInuse))
	metrics.PollCount.Add(1)

	// MSpanSys
	metrics.MSpanSys.Set(float64(memStats.MSpanSys))
	metrics.PollCount.Add(1)

	// Mallocs
	metrics.Mallocs.Set(float64(memStats.Mallocs))
	metrics.PollCount.Add(1)

	// NextGC
	metrics.NextGC.Set(float64(memStats.NextGC))
	metrics.PollCount.Add(1)

	// NumForcedGC
	metrics.NumForcedGC.Set(float64(memStats.NumForcedGC))
	metrics.PollCount.Add(1)

	// NumGC
	metrics.NumGC.Set(float64(memStats.NumGC))
	metrics.PollCount.Add(1)

	// OtherSys
	metrics.OtherSys.Set(float64(memStats.OtherSys))
	metrics.PollCount.Add(1)

	// PauseTotalNs
	metrics.PauseTotalNs.Set(float64(memStats.PauseTotalNs))
	metrics.PollCount.Add(1)

	// StackInuse
	metrics.StackInuse.Set(float64(memStats.StackInuse))
	metrics.PollCount.Add(1)

	// StackSys
	metrics.StackSys.Set(float64(memStats.StackSys))
	metrics.PollCount.Add(1)

	// Sys
	metrics.Sys.Set(float64(memStats.Sys))
	metrics.PollCount.Add(1)

	// TotalAlloc
	metrics.TotalAlloc.Set(float64(memStats.TotalAlloc))
	metrics.PollCount.Add(1)
}
