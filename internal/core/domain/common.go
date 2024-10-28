package domain

import (
	"errors"
	"fmt"
	"strings"
)

type MetricKind string

type MetricName string

type MetricID struct {
	kind MetricKind
	name MetricName
}

func (m *MetricID) Kind() MetricKind {
	return m.kind
}

func (m *MetricID) Name() MetricName {
	return m.name
}

func (m *MetricID) String() string {
	return fmt.Sprintf("%s/%s", m.kind, m.name)
}

func NewMetricID(kind MetricKind, name MetricName) *MetricID {
	return &MetricID{kind: kind, name: name}
}

type MetricValue interface {
	fmt.Stringer
}

type metricSpecificLogic interface {
	ParseValue(s string) (MetricValue, error)
	CalcNewValue(curr MetricValue, upd MetricValue) (MetricValue, error)
}

func ParseMetric(kind string, name string, value string) (*MetricID, MetricValue, error) {
	k, err := parseMetricKind(kind)
	if err != nil {
		return nil, nil, err
	}

	n, err := parseMetricName(name)
	if err != nil {
		return nil, nil, err
	}

	v, err := parseMetricValue(k, value)
	if err != nil {
		return nil, nil, err
	}

	return NewMetricID(k, n), v, nil
}

func CalcNewMetricValue(
	kind MetricKind, current MetricValue, update MetricValue,
) (MetricValue, error) {
	logic, ok := metrics[kind]
	if !ok {
		return nil, unknownKindErr(kind)
	}

	nv, err := logic.CalcNewValue(current, update)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return nv, nil
}

func parseMetricKind(s string) (MetricKind, error) {
	kind := MetricKind(s)
	if _, ok := metrics[kind]; !ok {
		return "", unknownKindErr(kind)
	}
	return kind, nil
}

func parseMetricName(s string) (MetricName, error) {
	if strings.TrimSpace(s) == "" {
		return "", invalidNameErr()
	}
	return MetricName(s), nil
}

func parseMetricValue(k MetricKind, s string) (MetricValue, error) {
	logic, ok := metrics[k]
	if !ok {
		return nil, unknownKindErr(k)
	}

	v, err := logic.ParseValue(s)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return v, nil
}

func castPairOfValues[T MetricValue](v1 MetricValue, v2 MetricValue) (T, T, error) {
	c1, ok1 := v1.(T)
	c2, ok2 := v2.(T)
	if !ok1 || !ok2 {
		var zero T
		return zero, zero, invalidValueErr(Counter, errors.New("values have wrong type"))
	}
	return c1, c2, nil
}
