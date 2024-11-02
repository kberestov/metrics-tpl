package domain

import (
	"errors"
	"fmt"
	"strconv"
)

// Counter is a kind of all counter metrics.
const Counter MetricKind = "counter"

// A CounterValue represents a specific type of counter metrics
// and implements the MetricValue interface.
type CounterValue int64

func (v CounterValue) String() string {
	return fmt.Sprintf("%d", v)
}

func (CounterValue) kind() MetricKind {
	return Counter
}

type counterLogicProvider struct{}

func (counterLogicProvider) parseValue(s string) (MetricValue, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, ErrInvalidMetricValue
	}

	return CounterValue(v), nil
}

func (counterLogicProvider) calcNewValue(c MetricValue, u MetricValue) (MetricValue, error) {
	ccv, ok := c.(CounterValue)
	if !ok {
		return nil, errors.New("current value is not CounterValue")
	}

	ucv, ok := u.(CounterValue)
	if !ok {
		return nil, errors.New("updating value is not CounterValue")
	}

	return ccv + ucv, nil
}
