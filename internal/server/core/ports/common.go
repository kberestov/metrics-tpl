package ports

import "errors"

var (
	ErrMetricNotFound      = errors.New("metric not found")
	ErrMetricValueRequired = errors.New("metric value required")
)
