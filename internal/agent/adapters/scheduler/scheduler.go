package scheduler

import (
	"sync"
	"time"

	"github.com/kberestov/metrics-tpl/internal/agent/config"
	"github.com/kberestov/metrics-tpl/internal/agent/core/ports"
)

type Scheduler struct {
	pollInterval   time.Duration
	reportInterval time.Duration
	poller         ports.MetricPoller
	reporter       ports.MetricReporter
}

func New(cfg config.Config, p ports.MetricPoller, r ports.MetricReporter) *Scheduler {
	return &Scheduler{
		poller:         p,
		pollInterval:   cfg.PollInterval,
		reporter:       r,
		reportInterval: cfg.ReportInterval,
	}
}

func (s *Scheduler) Run() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			s.poller.Poll()
			time.Sleep(s.pollInterval)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			time.Sleep(s.reportInterval)
			s.reporter.Report()
		}
	}()

	wg.Wait()
}
