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

func TestParseCounterValue(t *testing.T) {
	var zeroValue CounterValue
	tests := []struct {
		name    string
		arg     string
		want    CounterValue
		wantErr error
	}{
		{
			name:    "positive integer",
			arg:     "12345",
			want:    CounterValue(12345),
			wantErr: nil,
		},
		{
			name:    "negative integer",
			arg:     "-100",
			want:    CounterValue(-100),
			wantErr: nil,
		},
		{
			name:    "empty value",
			arg:     "",
			want:    zeroValue,
			wantErr: ErrInvalidMetricValue,
		},
		{
			name:    "float",
			arg:     "100.56",
			want:    zeroValue,
			wantErr: ErrInvalidMetricValue,
		},
		{
			name:    "not a number",
			arg:     "hello",
			want:    zeroValue,
			wantErr: ErrInvalidMetricValue,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := ParseCounterValue(test.arg)
			if test.wantErr == nil {
				require.NoError(t, err)
				assert.EqualValues(t, test.want, res)
				return
			}
			require.ErrorIs(t, err, test.wantErr)
		})
	}
}

func TestParseGaugeValue(t *testing.T) {
	var zeroValue GaugeValue
	tests := []struct {
		name    string
		arg     string
		want    GaugeValue
		wantErr error
	}{
		{
			name:    "integer",
			arg:     "12345",
			want:    GaugeValue(12345),
			wantErr: nil,
		},
		{
			name:    "positive float",
			arg:     "100.456",
			want:    GaugeValue(100.456),
			wantErr: nil,
		},
		{
			name:    "negative float",
			arg:     "-23.45",
			want:    GaugeValue(-23.45),
			wantErr: nil,
		},
		{
			name:    "empty",
			arg:     "",
			want:    zeroValue,
			wantErr: ErrInvalidMetricValue,
		},
		{
			name:    "empty",
			arg:     "",
			want:    zeroValue,
			wantErr: ErrInvalidMetricValue,
		},
		{
			name:    "not a number",
			arg:     "hello",
			want:    zeroValue,
			wantErr: ErrInvalidMetricValue,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := ParseGaugeValue(test.arg)
			if test.wantErr == nil {
				require.NoError(t, err)
				assert.EqualValues(t, test.want, res)
				return
			}
			require.ErrorIs(t, err, test.wantErr)
		})
	}
}
