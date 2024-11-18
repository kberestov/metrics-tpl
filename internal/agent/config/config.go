package config

import "time"

type Config struct {
	PollInterval   time.Duration
	ReportInterval time.Duration
	ServerREST     ServerREST
}

type ServerREST struct {
	Host   string
	Scheme string
}
