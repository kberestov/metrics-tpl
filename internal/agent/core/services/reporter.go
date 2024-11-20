package services

import (
	"fmt"
	"log"

	"github.com/kberestov/metrics-tpl/internal/agent/core/domain/metrics"
	"github.com/kberestov/metrics-tpl/internal/agent/core/ports"
	"github.com/kberestov/metrics-tpl/internal/common/domain"
)

var reportableMetrics = []metrics.Metric{
	metrics.PollCount,
	metrics.RandomValue,
	metrics.Alloc,
	metrics.BuckHashSys,
	metrics.Frees,
	metrics.GCCPUFraction,
	metrics.GCSys,
	metrics.HeapAlloc,
	metrics.HeapIdle,
	metrics.HeapInuse,
	metrics.HeapObjects,
	metrics.HeapReleased,
	metrics.HeapSys,
	metrics.LastGC,
	metrics.Lookups,
	metrics.MCacheInuse,
	metrics.MCacheSys,
	metrics.MSpanInuse,
	metrics.MSpanSys,
	metrics.Mallocs,
	metrics.NextGC,
	metrics.NumForcedGC,
	metrics.NumGC,
	metrics.OtherSys,
	metrics.PauseTotalNs,
	metrics.StackInuse,
	metrics.StackSys,
	metrics.Sys,
	metrics.TotalAlloc,
}

type MetricReporter struct {
	client ports.MetricServerClient
}

func NewMetricReporter(c ports.MetricServerClient) *MetricReporter {
	return &MetricReporter{client: c}
}

func (r *MetricReporter) Report() {
	for _, m := range reportableMetrics {
		go func() {
			// Remember the current value to perform correction
			// for counters after successful reporting.
			currentVal := m.Value()

			if err := r.client.UpdateMetric(m.Name(), currentVal); err != nil {
				log.Println(fmt.Errorf("failed to update metric: %w", err))
				return
			}

			// Make correction for counters.
			if currentVal.Kind() == domain.KindCounter {
				counter := m.(*metrics.Counter)
				correctingDelta := -currentVal.(domain.CounterValue)
				counter.Add(correctingDelta)
			}
		}()
	}
}
