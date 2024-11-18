package clients

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/kberestov/metrics-tpl/internal/agent/config"
	"github.com/kberestov/metrics-tpl/internal/common/domain"
)

type ServerRESTClient struct {
	client *http.Client
	cfg    config.ServerREST
}

func NewServerRESTClient(cfg config.Config) *ServerRESTClient {
	return &ServerRESTClient{
		cfg: cfg.ServerREST,
		client: &http.Client{
			// TODO: think about a proper value for the timeout.
			Timeout: cfg.ReportInterval,
		},
	}
}

// UpdateMetric sends a request to the following endpoint of the metric server:
// POST /update/{kind}/{name}/{value}.
func (c *ServerRESTClient) UpdateMetric(n domain.MetricName, v domain.MetricValue) error {
	if v == nil {
		return errors.New("no metric value provided")
	}

	u := url.URL{
		Scheme: c.cfg.Scheme,
		Host:   c.cfg.Host,
		Path: path.Join(
			"update",
			string(v.Kind()),
			string(n),
			v.String(),
		),
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), http.NoBody)
	if err != nil {
		return fmt.Errorf("failed to prepare request: %w", err)
	}

	req.Header.Set(`Content-Type`, `text/plain`)

	log.Printf("request: %s %s", req.Method, u.String())
	res, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	res.Body.Close()

	return nil
}
