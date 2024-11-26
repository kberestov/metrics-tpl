package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kberestov/metrics-tpl/internal/common/domain"
	"github.com/kberestov/metrics-tpl/internal/server/core/ports/mocks"
	"github.com/stretchr/testify/require"
)

func Test_handler_UpdateMetric(t *testing.T) {
	type pathValues struct {
		kind  string
		name  string
		value string
	}
	type updaterMockParams struct {
		times int
		n     domain.MetricName
		v     domain.MetricValue
		err   error
	}
	tests := []struct {
		name              string
		pathValues        pathValues
		updaterMockParams updaterMockParams
		wantCode          int
	}{
		{
			name:              "no path values",
			pathValues:        pathValues{},
			updaterMockParams: updaterMockParams{},
			wantCode:          http.StatusBadRequest,
		},
		{
			name: "unknown metric kind",
			pathValues: pathValues{
				kind: "unknown",
			},
			updaterMockParams: updaterMockParams{},
			wantCode:          http.StatusBadRequest,
		},
		{
			name: "counter with no name",
			pathValues: pathValues{
				kind: string(domain.KindCounter),
			},
			updaterMockParams: updaterMockParams{},
			wantCode:          http.StatusBadRequest,
		},
		{
			name: "counter with no value",
			pathValues: pathValues{
				kind: string(domain.KindCounter),
				name: "Counter",
			},
			updaterMockParams: updaterMockParams{},
			wantCode:          http.StatusBadRequest,
		},
		{
			name: "counter with gauge value",
			pathValues: pathValues{
				kind:  string(domain.KindCounter),
				name:  "Counter",
				value: domain.GaugeValue(10.56).String(),
			},
			updaterMockParams: updaterMockParams{},
			wantCode:          http.StatusBadRequest,
		},
		{
			name: "failed to update counter",
			pathValues: pathValues{
				kind:  string(domain.KindCounter),
				name:  "Counter",
				value: domain.CounterValue(10).String(),
			},
			updaterMockParams: updaterMockParams{
				times: 1,
				n:     domain.MetricName("Counter"),
				v:     domain.CounterValue(10),
				err:   errors.New("failed to update"),
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "successfully update gauge",
			pathValues: pathValues{
				kind:  string(domain.KindGauge),
				name:  "Gauge",
				value: domain.GaugeValue(10.56).String(),
			},
			updaterMockParams: updaterMockParams{
				times: 1,
				n:     domain.MetricName("Gauge"),
				v:     domain.GaugeValue(10.56),
				err:   nil,
			},
			wantCode: http.StatusOK,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/update", http.NoBody)
			req.SetPathValue(string(pvMetricKind), test.pathValues.kind)
			req.SetPathValue(string(pvMetricName), test.pathValues.name)
			req.SetPathValue(string(pvMetricValue), test.pathValues.value)

			w := httptest.NewRecorder()

			updaterMock := mocks.NewMetricUpdater(t)

			if test.updaterMockParams.times > 0 {
				updaterMock.
					On("Update", test.updaterMockParams.n, test.updaterMockParams.v).
					Return(test.updaterMockParams.err).
					Times(test.updaterMockParams.times)
			} else {
				updaterMock.AssertNotCalled(
					t,
					"Update",
					test.updaterMockParams.n,
					test.updaterMockParams.v,
				)
			}

			h := &handler{updater: updaterMock}

			h.UpdateMetric(w, req)

			res := w.Result()
			defer res.Body.Close()

			require.Equal(t, test.wantCode, res.StatusCode)
		})
	}
}
