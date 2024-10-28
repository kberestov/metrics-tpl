package domain

var metrics = map[MetricKind]metricSpecificLogic{
	Counter: counterLogic{},
	Gauge:   gaugeLogic{},
}
