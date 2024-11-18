package metrics

import (
	"sync/atomic"

	"github.com/kberestov/metrics-tpl/internal/common/domain"
)

type Metric interface {
	Name() domain.MetricName
	Value() domain.MetricValue
}

type Counter struct {
	n domain.MetricName
	v int64
}

func NewCounter(name string) *Counter {
	return &Counter{n: domain.MetricName(name)}
}

func (c *Counter) Name() domain.MetricName {
	return c.n
}

func (c *Counter) Value() domain.MetricValue {
	return domain.CounterValue(c.v)
}

func (c *Counter) Add(delta domain.CounterValue) {
	atomic.AddInt64(&c.v, int64(delta))
}

type Gauge struct {
	n domain.MetricName
	v float64
}

func NewGauge(name string) *Gauge {
	return &Gauge{n: domain.MetricName(name)}
}

func (g *Gauge) Name() domain.MetricName {
	return g.n
}

func (g *Gauge) Value() domain.MetricValue {
	return domain.GaugeValue(g.v)
}

func (g *Gauge) Set(value float64) {
	g.v = value
}

// Supported metrics:
var (
	PollCount     = NewCounter("PollCount")
	RandomValue   = NewGauge("RandomValue")
	Alloc         = NewGauge("Alloc")
	BuckHashSys   = NewGauge("BuckHashSys")
	Frees         = NewGauge("Frees")
	GCCPUFraction = NewGauge("GCCPUFraction")
	GCSys         = NewGauge("GCSys")
	HeapAlloc     = NewGauge("HeapAlloc")
	HeapIdle      = NewGauge("HeapIdle")
	HeapInuse     = NewGauge("HeapInuse")
	HeapObjects   = NewGauge("HeapObjects")
	HeapReleased  = NewGauge("HeapReleased")
	HeapSys       = NewGauge("HeapSys")
	LastGC        = NewGauge("LastGC")
	Lookups       = NewGauge("Lookups")
	MCacheInuse   = NewGauge("MCacheInuse")
	MCacheSys     = NewGauge("MCacheSys")
	MSpanInuse    = NewGauge("MSpanInuse")
	MSpanSys      = NewGauge("MSpanSys")
	Mallocs       = NewGauge("Mallocs")
	NextGC        = NewGauge("NextGC")
	NumForcedGC   = NewGauge("NumForcedGC")
	NumGC         = NewGauge("NumGC")
	OtherSys      = NewGauge("OtherSys")
	PauseTotalNs  = NewGauge("PauseTotalNs")
	StackInuse    = NewGauge("StackInuse")
	StackSys      = NewGauge("StackSys")
	Sys           = NewGauge("Sys")
	TotalAlloc    = NewGauge("TotalAlloc")
)
