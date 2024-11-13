package ports

// MetricService represents use cases related to metrics.
type MetricService interface {
	// Update updates a storable value of a metric
	// of the given kind and name with the given value.
	Update(kind string, name string, value string) error
}
