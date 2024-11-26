package services

import (
	"errors"
	"testing"

	"github.com/kberestov/metrics-tpl/internal/common/domain"
	dmocks "github.com/kberestov/metrics-tpl/internal/common/domain/mocks"
	"github.com/kberestov/metrics-tpl/internal/server/core/ports"
	pmocks "github.com/kberestov/metrics-tpl/internal/server/core/ports/mocks"
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

func TestMetricUpdater_Update(t *testing.T) {
	type args struct {
		n domain.MetricName
		v domain.MetricValue
	}
	type getValueParams struct {
		times int
		val   domain.MetricValue
		err   error
	}
	type saveValueParams struct {
		times  int
		newVal domain.MetricValue
		err    error
	}
	type storeMockParams struct {
		getValue  getValueParams
		saveValue saveValueParams
	}

	unknownMetricVal := dmocks.NewMetricValue(t)
	unknownMetricVal.On("Kind").Return(domain.MetricKind("unknown"))

	tests := []struct {
		name            string
		args            args
		storeMockParams storeMockParams
		wantErr         bool
	}{
		{
			name: "no metric value",
			args: args{domain.MetricName("m1"), nil},
			storeMockParams: storeMockParams{
				getValue: getValueParams{
					times: 0,
					val:   nil,
					err:   nil,
				},
				saveValue: saveValueParams{
					times:  0,
					newVal: nil,
					err:    nil,
				},
			},
			wantErr: true,
		},
		{
			name: "unknown metric value kind",
			args: args{domain.MetricName("m1"), unknownMetricVal},
			storeMockParams: storeMockParams{
				getValue: getValueParams{
					times: 0,
					val:   nil,
					err:   nil,
				},
				saveValue: saveValueParams{
					times:  0,
					newVal: nil,
					err:    nil,
				},
			},
			wantErr: true,
		},
		{
			name: "counter - failed to get current value",
			args: args{domain.MetricName("Counter 1"), domain.CounterValue(100)},
			storeMockParams: storeMockParams{
				getValue: getValueParams{
					times: 1,
					val:   nil,
					err:   errors.New("some error occured"),
				},
				saveValue: saveValueParams{
					times:  0,
					newVal: nil,
					err:    nil,
				},
			},
			wantErr: true,
		},
		{
			name: "counter - current value mismatch",
			args: args{domain.MetricName("Counter 1"), domain.CounterValue(100)},
			storeMockParams: storeMockParams{
				getValue: getValueParams{
					times: 1,
					val:   domain.GaugeValue(300.45),
					err:   nil,
				},
				saveValue: saveValueParams{
					times:  0,
					newVal: nil,
					err:    nil,
				},
			},
			wantErr: true,
		},
		{
			name: "counter - failed to save a new value",
			args: args{domain.MetricName("Counter 1"), domain.CounterValue(100)},
			storeMockParams: storeMockParams{
				getValue: getValueParams{
					times: 1,
					val:   nil,
					err:   ports.ErrMetricNotFound,
				},
				saveValue: saveValueParams{
					times:  1,
					newVal: domain.CounterValue(100),
					err:    errors.New("some error occured"),
				},
			},
			wantErr: true,
		},
		{
			name: "counter - saving a new value",
			args: args{domain.MetricName("Counter 1"), domain.CounterValue(100)},
			storeMockParams: storeMockParams{
				getValue: getValueParams{
					times: 1,
					val:   nil,
					err:   ports.ErrMetricNotFound,
				},
				saveValue: saveValueParams{
					times:  1,
					newVal: domain.CounterValue(100),
					err:    nil,
				},
			},
			wantErr: false,
		},
		{
			name: "counter - saving an updated value",
			args: args{domain.MetricName("Counter 1"), domain.CounterValue(100)},
			storeMockParams: storeMockParams{
				getValue: getValueParams{
					times: 1,
					val:   domain.CounterValue(50),
					err:   nil,
				},
				saveValue: saveValueParams{
					times:  1,
					newVal: domain.CounterValue(150),
					err:    nil,
				},
			},
			wantErr: false,
		},
		{
			name: "gauge - failed to save an updated value",
			args: args{domain.MetricName("Gauge 1"), domain.GaugeValue(50.56)},
			storeMockParams: storeMockParams{
				getValue: getValueParams{
					times: 0,
					val:   nil,
					err:   nil,
				},
				saveValue: saveValueParams{
					times:  1,
					newVal: domain.GaugeValue(50.56),
					err:    errors.New("some error occured"),
				},
			},
			wantErr: true,
		},
		{
			name: "gauge - saving an updated value",
			args: args{domain.MetricName("Gauge 1"), domain.GaugeValue(50.56)},
			storeMockParams: storeMockParams{
				getValue: getValueParams{
					times: 0,
					val:   nil,
					err:   nil,
				},
				saveValue: saveValueParams{
					times:  1,
					newVal: domain.GaugeValue(50.56),
					err:    nil,
				},
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storeMock := pmocks.NewMetricStore(t)

			if test.storeMockParams.getValue.times > 0 {
				storeMock.
					On("GetValue", test.args.n).
					Return(test.storeMockParams.getValue.val, test.storeMockParams.getValue.err).
					Times(test.storeMockParams.getValue.times)
			} else {
				storeMock.AssertNotCalled(t, "GetValue", test.args.n)
			}

			if test.storeMockParams.saveValue.times > 0 {
				storeMock.
					On("SaveValue", test.args.n, test.storeMockParams.saveValue.newVal).
					Return(test.storeMockParams.saveValue.err).
					Times(test.storeMockParams.saveValue.times)
			} else {
				storeMock.AssertNotCalled(
					t,
					"SaveValue",
					test.args.n,
					test.storeMockParams.saveValue.newVal,
				)
			}

			updater := &MetricUpdater{store: storeMock}

			err := updater.Update(test.args.n, test.args.v)

			if test.wantErr {
				require.Error(t, err)
			}
		})
	}
}
