package domain

import (
	"fmt"
	"strconv"
)

const Counter MetricKind = "counter"

type CounterValue int64

func (v *CounterValue) String() string {
	return fmt.Sprintf("%d", *v)
}

type counterLogic struct{}

func (counterLogic) ParseValue(s string) (MetricValue, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, invalidValueErr(Counter, err)
	}
	cv := CounterValue(v)
	return &cv, nil
}

func (counterLogic) CalcNewValue(curr MetricValue, upd MetricValue) (MetricValue, error) {
	cv, uv, err := castPairOfValues[*CounterValue](curr, upd)
	if err != nil {
		return nil, err
	}
	nv := *cv + *uv
	return &nv, nil
}
