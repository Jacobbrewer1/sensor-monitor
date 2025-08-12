package main

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	mockalerts "github.com/jacobbrewer1/sensor-monitor/pkg/alerts/mock"
)

func TestWatchCPUCoreTemperatures(t *testing.T) {
	t.Parallel()

	t.Run("golden_error_handling", func(t *testing.T) {
		t.Parallel()

		// Mock logger and alerter
		logger := slog.New(slog.DiscardHandler)

		mockController := gomock.NewController(t)
		alerter := mockalerts.NewMockAlerter(mockController)

		alerter.EXPECT().Alert(
			"Reading CPU Core Temperatures Failed",
			"An error occurred while reading CPU core temperatures: test error",
		).Times(1)

		// Create channels for temperatures and errors
		tempChan := make(chan map[int]float64)
		errChan := make(chan error)

		t.Cleanup(func() {
			close(tempChan)
			close(errChan)
		})

		// Start the watchCPUCoreTemperatures function in a goroutine
		go watchCPUCoreTemperatures(
			t.Context(),
			logger,
			alerter,
			tempChan,
			errChan,
		)

		// Send an error to the error channel
		err := errors.New("test error")
		errChan <- err
	})
}

func TestTempWithBuffer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		inputTemp    float64
		inputPercent int
		expected     float64
	}{
		{
			name:         "normal_temperature",
			inputTemp:    50.0,
			inputPercent: 0,
			expected:     50.0,
		},
		{
			name:         "high_temperature",
			inputTemp:    90.0,
			inputPercent: 10,
			expected:     99.0,
		},
		{
			name:         "extreme_temperature",
			inputTemp:    100.0,
			inputPercent: 20,
			expected:     120.0,
		},
		{
			name:         "negative_temperature",
			inputTemp:    -10.0,
			inputPercent: 0,
			expected:     -10.0,
		},
		{
			name:         "zero_temperature",
			inputTemp:    0.0,
			inputPercent: 0,
			expected:     0.0,
		},
		{
			name:         "negative_percent",
			inputTemp:    50.0,
			inputPercent: -10,
			expected:     45.0, // Assuming negative percent reduces the temperature
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := tempWithBuffer(tt.inputTemp, tt.inputPercent)
			require.Equal(t, tt.expected, result)
		})
	}
}
