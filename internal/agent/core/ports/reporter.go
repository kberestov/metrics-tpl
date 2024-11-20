package ports

// MetricReporter represents an application service for
// reporting about metrics collected by the agent.
type MetricReporter interface {
	// Report reports once about current state of all the metrics.
	Report()
}
