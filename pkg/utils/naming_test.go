package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrettyName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", "Unknown"},
		{"Single word", "sensor", "Sensor"},
		{"Underscores", "sensor_name", "Sensor Name"},
		{"Mixed case", "sensorName", "Sensor Name"},
		{"Multiple underscores", "sensor_name_test", "Sensor Name Test"},
		{"Multiple underscores", "sensor-name-test", "Sensor Name Test"},
		{"Leading and trailing spaces", "  sensor_name  ", "Sensor Name"},
		{"Already pretty", "Sensor Name", "Sensor Name"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := PrettyName(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
