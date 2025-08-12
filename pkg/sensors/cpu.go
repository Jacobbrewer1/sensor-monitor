package sensors

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// WatchCPUCoreTemperature returns a channel that emits CPU core temperatures
// and a channel for errors. It continuously monitors CPU core temperatures.
func WatchCPUCoreTemperature(watchInterval time.Duration) (<-chan map[int]float64, <-chan error) {
	return WatchCPUCoreTemperatureWithContext(context.Background(), watchInterval)
}

// WatchCPUCoreTemperatureWithContext returns channels for CPU core temperatures and errors,
// with support for context cancellation. It monitors temperatures every 1 second to avoid CPU throttling.
func WatchCPUCoreTemperatureWithContext(ctx context.Context, watchInterval time.Duration) (<-chan map[int]float64, <-chan error) {
	tempChan := make(chan map[int]float64)
	errChan := make(chan error)

	go func() {
		defer close(tempChan)
		defer close(errChan)

		// Send initial reading
		if temps, err := CPUCoreTemperatures(); err != nil {
			publishNonBlockingChannel(errChan, err)
		} else {
			publishNonBlockingChannel(tempChan, temps)
		}

		ticker := time.NewTicker(watchInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				temps, err := CPUCoreTemperatures()
				if err != nil {
					publishNonBlockingChannel(errChan, err)
					continue // Skip sending data if there's an error
				}

				publishNonBlockingChannel(tempChan, temps)
			}
		}
	}()

	return tempChan, errChan
}

func publishNonBlockingChannel[T any](ch chan<- T, data T) {
	select {
	case ch <- data:
	default:
		// Channel is full, do nothing
	}
}

// CPUCoreTemperatures returns a map of CPU core ID to temperature in Celsius
func CPUCoreTemperatures() (map[int]float64, error) {
	switch runtime.GOOS {
	case "linux":
		return cpuTemperatureLinux()
	default:
		return nil, fmt.Errorf("CPU temperature monitoring not supported on %s", runtime.GOOS)
	}
}

func cpuTemperatureLinux() (map[int]float64, error) {
	temps := make(map[int]float64)

	// Look for hwmon devices
	hwmonDevices, err := filepath.Glob("/sys/class/hwmon/hwmon*")
	if err != nil {
		return nil, fmt.Errorf("failed to find hwmon devices: %w", err)
	}

	if len(hwmonDevices) == 0 {
		return nil, fmt.Errorf("no hwmon devices found")
	}

	coreIndex := 0
	for _, hwmonPath := range hwmonDevices {
		// Check if this hwmon device is CPU-related by reading the name
		namePath := filepath.Join(hwmonPath, "name")
		nameData, err := os.ReadFile(namePath)
		if err != nil {
			continue // Skip if we can't read the name
		}

		hwmonName := strings.TrimSpace(string(nameData))

		// Check if this hwmon device is CPU-related
		if !isCPUHwmon(hwmonName) {
			continue
		}

		// Look for temperature inputs in this hwmon device
		tempInputs, err := filepath.Glob(filepath.Join(hwmonPath, "temp*_input"))
		if err != nil {
			continue // Skip if we can't find temp inputs
		}

		for _, tempInput := range tempInputs {
			// Read temperature from this input
			tempData, err := os.ReadFile(tempInput)
			if err != nil {
				continue // Skip if we can't read temperature
			}

			tempStr := strings.TrimSpace(string(tempData))
			tempMilliC, err := strconv.ParseInt(tempStr, 10, 64)
			if err != nil {
				continue // Skip if we can't parse temperature
			}

			// Convert from millicelsius to celsius
			tempC := float64(tempMilliC) / 1000.0

			// Skip unreasonable temperatures (likely sensors that aren't working)
			if tempC < 0 || tempC > 150 {
				continue
			}

			temps[coreIndex] = tempC
			coreIndex++
		}
	}

	if len(temps) == 0 {
		return nil, fmt.Errorf("no CPU temperature sensors found")
	}

	return temps, nil
}

// isCPUHwmon checks if a hwmon device name corresponds to a CPU sensor
func isCPUHwmon(hwmonName string) bool {
	cpuHwmonNames := []string{
		"coretemp",
		"k10temp",     // AMD Ryzen/Threadripper
		"k8temp",      // AMD K8
		"via-cputemp", // VIA
		"cpu_thermal", // ARM/embedded
		"soc_thermal", // System on Chip
		"acpi",        // ACPI thermal (sometimes CPU)
	}

	hwmonNameLower := strings.ToLower(hwmonName)
	for _, cpuName := range cpuHwmonNames {
		if strings.Contains(hwmonNameLower, strings.ToLower(cpuName)) {
			return true
		}
	}

	return false
}
