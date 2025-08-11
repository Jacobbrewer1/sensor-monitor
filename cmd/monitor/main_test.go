package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldNotify(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		currentTemp float64
		lastTemp    float64
		expected    bool
	}{
		{
			name:        "no previous temperature (lastTemp is 0)",
			currentTemp: 90.0,
			lastTemp:    0.0,
			expected:    false,
		},
		{
			name:        "current temp below worry threshold (85% of crash temp)",
			currentTemp: 80.0,
			lastTemp:    75.0,
			expected:    false,
		},
		{
			name:        "current temp at worry threshold but no significant increase",
			currentTemp: 85.0,
			lastTemp:    85.0,
			expected:    false,
		},
		{
			name:        "current temp above worry threshold with 5% increase",
			currentTemp: 89.26, // 85 * 1.05 = 89.25, so 89.26 > 89.25
			lastTemp:    85.0,
			expected:    true,
		},
		{
			name:        "current temp above worry threshold with slight increase (under 5%)",
			currentTemp: 87.0,
			lastTemp:    85.0,
			expected:    false,
		},
		{
			name:        "current temp at crash temperature",
			currentTemp: 100.0,
			lastTemp:    90.0,
			expected:    true,
		},
		{
			name:        "current temp above crash temperature",
			currentTemp: 105.0,
			lastTemp:    90.0,
			expected:    true,
		},
		{
			name:        "current temp at crash temp but below last temp",
			currentTemp: 100.0,
			lastTemp:    110.0,
			expected:    true,
		},
		{
			name:        "edge case: exactly at worry threshold",
			currentTemp: 85.0,
			lastTemp:    80.0,
			expected:    true, // 80 * 1.05 = 84, so 85 > 84 and 85 >= 85 (worry threshold)
		},
		{
			name:        "edge case: just above 5% threshold at worry level",
			currentTemp: 84.1, // 80 * 1.05 = 84, so 84.1 > 84
			lastTemp:    80.0,
			expected:    false, // below worry threshold (85)
		},
		{
			name:        "large temperature drop but still above worry threshold",
			currentTemp: 90.0,
			lastTemp:    95.0,
			expected:    false, // 95 * 1.05 = 99.75, 90 < 99.75 and 90 < 100
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := shouldNotify(test.currentTemp, test.lastTemp, crashTemperature)
			require.Equal(t, test.expected, result)
		})
	}
}
