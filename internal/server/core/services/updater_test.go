package services

import (
	"testing"

	"github.com/kberestov/metrics-tpl/internal/common/domain"
	"github.com/kberestov/metrics-tpl/internal/server/core/ports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_calcNewCounterValue(t *testing.T) {
	type args struct {
		curr domain.MetricValue
		upd  domain.CounterValue
	}
	tests := []struct {
		name    string
		args    args
		want    domain.MetricValue
		wantErr error
	}{
		{
			name:    "no current value",
			args:    args{nil, domain.CounterValue(100)},
			want:    domain.CounterValue(100),
			wantErr: nil,
		},
		{
			name:    "updating gauge with counter",
			args:    args{domain.GaugeValue(20.5), domain.CounterValue(10)},
			want:    nil,
			wantErr: ports.ErrMetricValueKindMismatch,
		},
		{
			name:    "new = current + updating",
			args:    args{domain.CounterValue(30), domain.CounterValue(10)},
			want:    domain.CounterValue(40),
			wantErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := calcNewCounterValue(test.args.curr, test.args.upd)
			if test.wantErr == nil {
				require.NoError(t, err)
				assert.EqualValues(t, test.want, got)
				return
			}
			require.ErrorIs(t, err, test.wantErr)
		})
	}
}
