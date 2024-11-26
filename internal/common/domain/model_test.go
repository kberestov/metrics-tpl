package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMetricKind(t *testing.T) {
	var zeroKind MetricKind
	tests := []struct {
		name    string
		arg     string
		want    MetricKind
		wantErr error
	}{
		{
			name:    "valid counter kind",
			arg:     "counter",
			want:    KindCounter,
			wantErr: nil,
		},
		{
			name:    "invalid counter kind",
			arg:     "Counter",
			want:    zeroKind,
			wantErr: ErrUnknownMetricKind,
		},
		{
			name:    "valid gauge kind",
			arg:     "gauge",
			want:    KindGauge,
			wantErr: nil,
		},
		{
			name:    "unknown kind",
			arg:     "abcd",
			want:    zeroKind,
			wantErr: ErrUnknownMetricKind,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := ParseMetricKind(test.arg)
			if test.wantErr == nil {
				require.NoError(t, err)
				assert.EqualValues(t, test.want, res)
				return
			}
			require.ErrorIs(t, err, test.wantErr)
		})
	}
}

func TestParseMetricName(t *testing.T) {
	var zeroName MetricName
	tests := []struct {
		name    string
		arg     string
		want    MetricName
		wantErr error
	}{
		{
			name:    "valid name 1",
			arg:     "counter 1",
			want:    MetricName("counter 1"),
			wantErr: nil,
		},
		{
			name:    "valid name 2",
			arg:     " Alloc   ",
			want:    MetricName(" Alloc   "),
			wantErr: nil,
		},
		{
			name:    "empty name",
			arg:     "",
			want:    zeroName,
			wantErr: ErrInvalidMetricName,
		},
		{
			name:    "whitespace name",
			arg:     "  ",
			want:    zeroName,
			wantErr: ErrInvalidMetricName,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := ParseMetricName(test.arg)
			if test.wantErr == nil {
				require.NoError(t, err)
				assert.EqualValues(t, test.want, res)
				return
			}
			require.ErrorIs(t, err, test.wantErr)
		})
	}
}

func TestParseMetricValue(t *testing.T) {
	type args struct {
		k MetricKind
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    MetricValue
		wantErr error
	}{
		{
			name:    "valid counter",
			args:    args{KindCounter, "123"},
			want:    CounterValue(123),
			wantErr: nil,
		},
		{
			name:    "valid gauge",
			args:    args{KindGauge, "123.56"},
			want:    GaugeValue(123.56),
			wantErr: nil,
		},
		{
			name:    "float counter",
			args:    args{KindCounter, "123.56"},
			want:    nil,
			wantErr: ErrInvalidMetricValue,
		},
		{
			name:    "empty counter value",
			args:    args{KindCounter, ""},
			want:    nil,
			wantErr: ErrInvalidMetricValue,
		},
		{
			name:    "empty gauge value",
			args:    args{KindGauge, ""},
			want:    nil,
			wantErr: ErrInvalidMetricValue,
		},
		{
			name:    "not a number",
			args:    args{KindCounter, "abcd"},
			want:    nil,
			wantErr: ErrInvalidMetricValue,
		},
		{
			name:    "unknown kind",
			args:    args{MetricKind("unknown"), "abcd"},
			want:    nil,
			wantErr: ErrUnknownMetricKind,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := ParseMetricValue(test.args.k, test.args.s)
			if test.wantErr == nil {
				require.NoError(t, err)
				assert.EqualValues(t, test.want, res)
				return
			}
			require.ErrorIs(t, err, test.wantErr)
		})
	}
}

func TestCounterValue_Kind(t *testing.T) {
	var v CounterValue
	require.EqualValues(t, KindCounter, v.Kind())
}

func TestGaugeValue_Kind(t *testing.T) {
	var v GaugeValue
	require.EqualValues(t, KindGauge, v.Kind())
}
