package domain

// metricRegistry keeps all supported kinds of metrics
// along with their specific logic.
var metricRegistry = map[MetricKind]metricLogic{
	Counter: counterLogicProvider{},
	Gauge:   gaugeLogicProvider{},
}
