package ports

import "errors"

var (
	ErrMetricNotFound          = errors.New("metric not found")
	ErrNoMetricValue           = errors.New("metric value not present")
	ErrMetricValueKindMismatch = errors.New("found metric has mismatched type")
)
