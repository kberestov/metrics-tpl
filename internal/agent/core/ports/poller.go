package ports

// MetricPoller represents an application service to poll metrics
// collected by the agent.
type MetricPoller interface {
	// Poll updates once the values of all the metrics.
	Poll()
}
