package domain

import (
	"errors"
	"fmt"
	"strings"
)

// A MetricKind represents a domain type for metric kinds.
// Each supported metric must have own unique value of MetricKind.
type MetricKind string

// A MetricName represents a domain type for metric names.
type MetricName string

// A MetricID represents a unit to identify a specific metric:
// a pair of MetricKind and MetricName.
type MetricID struct {
	k MetricKind
	n MetricName
}

func (id MetricID) Kind() MetricKind {
	return id.k
}

func (id MetricID) Name() MetricName {
	return id.n
}

// NewMetricID creates a new MetricID.
func NewMetricID(kind MetricKind, name MetricName) MetricID {
	return MetricID{k: kind, n: name}
}

// MetricValue represents any metric type values depending on metric kind.
// Each supported metric must provide own impelementation of the interface.
type MetricValue interface {
	fmt.Stringer
	kind() MetricKind
}

// metricLogic represents a specific logic that differs depending on metric kind.
// Each supported metric must provide own impelementation of the interface.
type metricLogic interface {
	// parseValue tries to parse the raw string value to create a valid MetricValue.
	parseValue(s string) (MetricValue, error)

	// calcNewValue tries to calculate a new metric value during the metric update:
	// it applies the updating value to the current value.
	// Both current and updating values are not nil.
	calcNewValue(c MetricValue, u MetricValue) (MetricValue, error)
}

var (
	ErrUnknownMetricKind  = errors.New("unknown metric kind")
	ErrInvalidMetricName  = errors.New("invalid metric name")
	ErrInvalidMetricValue = errors.New("invalid metric value")
)

// ParseMetricID tries to parse the raw string values to create a valid MetricID.
func ParseMetricID(kind string, name string) (MetricID, error) {
	k, err := parseMetricKind(kind)
	if err != nil {
		var zero MetricID
		return zero, err
	}

	n, err := parseMetricName(name)
	if err != nil {
		var zero MetricID
		return zero, err
	}

	return NewMetricID(k, n), nil
}

// ParseMetricValue tries to parse the raw string value to create a valid MetricValue
// depending on the given metric kind.
func ParseMetricValue(kind MetricKind, s string) (MetricValue, error) {
	logic, ok := metricRegistry[kind]
	if !ok {
		return nil, ErrUnknownMetricKind
	}

	v, err := logic.parseValue(s)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return v, nil
}

// CalcNewMetricValue tries to calculate a new metric value during the metric update.
// It applies the updating value to the current value depending on the given metric kind.
// Edge cases:
//   - if the current value is nil, the updating will be returned
//   - if the updating value is nil, the current will be returned
func CalcNewMetricValue(kind MetricKind, current MetricValue, updating MetricValue) (
	newValue MetricValue,
	err error,
) {
	logic, ok := metricRegistry[kind]
	if !ok {
		return nil, ErrUnknownMetricKind
	}

	if current == nil {
		return updating, nil
	}

	if updating == nil {
		return current, nil
	}

	newValue, err = logic.calcNewValue(current, updating)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return newValue, nil
}

func parseMetricKind(s string) (MetricKind, error) {
	kind := MetricKind(s)
	if _, ok := metricRegistry[kind]; !ok {
		return "", ErrUnknownMetricKind
	}
	return kind, nil
}

func parseMetricName(s string) (MetricName, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", ErrInvalidMetricName
	}
	return MetricName(s), nil
}
