package domain

import (
	"errors"
	"fmt"
	"strings"
)

type MetricKind string

const (
	Gauge   MetricKind = "gauge"
	Counter MetricKind = "counter"
)

type MetricName string

type Metric interface {
	Kind() MetricKind
	Name() MetricName
}

type metric struct {
	kind MetricKind
	name MetricName
}

func (m metric) Kind() MetricKind {
	return m.kind
}

func (m metric) Name() MetricName {
	return m.name
}

type (
	GaugeValue   float64
	CounterValue int64
)

func ParseMetric(kind string, name string) (Metric, error) {
	k, err := parseMetricKind(kind)
	if err != nil {
		return nil, err
	}

	n, err := parseMetricName(name)
	if err != nil {
		return nil, err
	}

	return metric{kind: k, name: n}, nil
}

func parseMetricKind(s string) (MetricKind, error) {
	kind := MetricKind(s)
	switch kind {
	case Gauge, Counter:
		return kind, nil
	default:
		return "", fmt.Errorf("unknown metric kind: %s", s)
	}
}

func parseMetricName(s string) (MetricName, error) {
	if strings.TrimSpace(s) == "" {
		return "", errors.New("invalid metric name")
	}
	return MetricName(s), nil
}
