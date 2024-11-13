package domain

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrUnknownMetricKind  = errors.New("unknown metric kind")
	ErrInvalidMetricName  = errors.New("invalid metric name")
	ErrInvalidMetricValue = errors.New("invalid metric value")
)

type MetricKind string

const (
	KindCounter MetricKind = "counter"
	KindGauge   MetricKind = "gauge"
)

func ParseMetricKind(s string) (MetricKind, error) {
	k := MetricKind(s)
	switch k {
	case KindCounter, KindGauge:
		return k, nil
	default:
		var zero MetricKind
		return zero, ErrUnknownMetricKind
	}
}

type MetricName string

func ParseMetricName(s string) (MetricName, error) {
	if strings.TrimSpace(s) == "" {
		var zero MetricName
		return zero, ErrInvalidMetricName
	}
	return MetricName(s), nil
}

type (
	CounterValue int64
	GaugeValue   float64
)

func ParseCounterValue(s string) (CounterValue, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		var zero CounterValue
		return zero, ErrInvalidMetricValue
	}
	return CounterValue(v), nil
}

func ParseGaugeValue(s string) (GaugeValue, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		var zero GaugeValue
		return zero, ErrInvalidMetricValue
	}
	return GaugeValue(v), nil
}
