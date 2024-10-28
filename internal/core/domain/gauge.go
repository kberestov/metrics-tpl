package domain

import (
	"fmt"
	"strconv"
)

const Gauge MetricKind = "gauge"

type GaugeValue float64

func (v *GaugeValue) String() string {
	return fmt.Sprintf("%f", *v)
}

type gaugeLogic struct{}

func (gaugeLogic) ParseValue(s string) (MetricValue, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, invalidValueErr(Gauge, err)
	}
	gv := GaugeValue(v)
	return &gv, nil
}

func (gaugeLogic) CalcNewValue(curr MetricValue, upd MetricValue) (MetricValue, error) {
	_, uv, err := castPairOfValues[*GaugeValue](curr, upd)
	if err != nil {
		return nil, err
	}
	nv := *uv
	return &nv, nil
}
