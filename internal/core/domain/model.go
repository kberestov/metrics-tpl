package domain

import (
	"errors"
	"fmt"
	"strconv"
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
	ParseValue(s string) (MetricValue, error)
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

func (m metric) ParseValue(s string) (MetricValue, error) {
	invalidValueErr := func(err error) error {
		return fmt.Errorf("invalid value for %v kind of metric: %w", m.kind, err)
	}

	switch m.kind {
	case Counter:
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, invalidValueErr(err)
		}
		return CounterValue{Value: CounterValueType(v), metric: m}, nil
	case Gauge:
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, invalidValueErr(err)
		}
		return CounterValue{Value: CounterValueType(v), metric: m}, nil
	default:
		return nil, fmt.Errorf("unknown metric kind: %v", m.kind)
	}
}

func (m metric) Metric() Metric {
	return m
}

type MetricValue interface {
	Metric() Metric
	fmt.Stringer
}

type (
	GaugeValueType   float64
	CounterValueType int64
)

type GaugeValue struct {
	metric
	Value GaugeValueType
}

func (gv GaugeValue) String() string {
	return fmt.Sprintf("%v", gv.Value)
}

type CounterValue struct {
	metric
	Value CounterValueType
}

func (cv CounterValue) String() string {
	return fmt.Sprintf("%v", cv.Value)
}

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
