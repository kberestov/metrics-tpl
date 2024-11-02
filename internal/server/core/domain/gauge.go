package domain

import (
	"errors"
	"fmt"
	"strconv"
)

// Gauge is a kind of all gauge metrics.
const Gauge MetricKind = "gauge"

// A GaugeValue represents a specific type of gauge metrics
// and implements the MetricValue interface.
type GaugeValue float64

func (v GaugeValue) String() string {
	return fmt.Sprintf("%f", v)
}

func (GaugeValue) kind() MetricKind {
	return Gauge
}

type gaugeLogicProvider struct{}

func (gaugeLogicProvider) parseValue(s string) (MetricValue, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, ErrInvalidMetricValue
	}

	return GaugeValue(v), nil
}

func (gaugeLogicProvider) calcNewValue(c MetricValue, u MetricValue) (MetricValue, error) {
	ugv, ok := u.(GaugeValue)
	if !ok {
		return nil, errors.New("updating value is not GaugeValue")
	}

	return ugv, nil
}
