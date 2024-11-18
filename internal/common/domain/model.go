package domain

import (
	"errors"
	"fmt"
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

type MetricValue interface {
	fmt.Stringer
	Kind() MetricKind
}

type (
	CounterValue int64
	GaugeValue   float64
)

func (v CounterValue) Kind() MetricKind {
	return KindCounter
}

func (v CounterValue) String() string {
	return fmt.Sprintf("%d", v)
}

func (v GaugeValue) Kind() MetricKind {
	return KindGauge
}

func (v GaugeValue) String() string {
	return fmt.Sprintf("%f", v)
}

func ParseMetricValue(k MetricKind, s string) (MetricValue, error) {
	switch k {
	case KindCounter:
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			var zero CounterValue
			return zero, ErrInvalidMetricValue
		}
		return CounterValue(v), nil
	case KindGauge:
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			var zero GaugeValue
			return zero, ErrInvalidMetricValue
		}
		return GaugeValue(v), nil
	default:
		return nil, ErrUnknownMetricKind
	}
}
