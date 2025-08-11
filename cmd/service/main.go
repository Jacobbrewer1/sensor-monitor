package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

const (
	// appName is the name of the application used for notifications.
	appName = "CPU Temp Monitor"

	// crashTemperature is the temperature in Celsius at which the system is considered to be in a critical state.
	crashTemperature = 100.0 // Temperature in Celsius at which the system crashes
)

func readCPUTemp() (float64, error) {
	cmd := exec.Command("sensors", "-j")
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("failed to execute sensors command: %w", err)
	}

	sensorData := new(sensor)
	if err := json.NewDecoder(strings.NewReader(string(output))).Decode(sensorData); err != nil {
		return 0, fmt.Errorf("failed to decode sensors output: %w", err)
	}

	return sensorData.DellDdvVirtual0.CPU.Temp1Input, nil
}

func notifyUser(currentTemp float64) error {
	if currentTemp >= crashTemperature {
		if err := beeep.Notify(
			"🔥 CPU Temperature Critical!",
			fmt.Sprintf("CPU has reached %.1f°C — system will crash soon!", currentTemp),
			"",
		); err != nil {
			return fmt.Errorf("failed to send critical notification: %w", err)
		}

		return nil
	}

	if err := beeep.Notify(
		"⚠ CPU Temperature Alert",
		fmt.Sprintf("CPU temperature is at %.1f°C — please check your system!", currentTemp),
		"",
	); err != nil {
		return fmt.Errorf("failed to send beep notification: %w", err)
	}

	return nil
}

func shouldNotify(currentTemp, lastTemp float64) bool {
	if lastTemp == 0 {
		return false // No previous temperature to compare
	}

	crashWorryThreshold := crashTemperature * 0.85 // 85% of crash temperature
	if currentTemp < crashWorryThreshold {
		return false // Current temperature is below the threshold for concern
	}

	// Calculate the threshold for notification
	threshold := lastTemp * 1.05 // 5% increase from the last temperature

	// Notify if the current temperature is significantly higher than the last recorded temperature
	return currentTemp > threshold || currentTemp >= crashTemperature
}

func main() {
	var lastTemp float64
	beeep.AppName = appName

	for {
		temp, err := readCPUTemp()
		if err != nil {
			fmt.Printf("Error reading CPU temperature: %v\n", err)
			return
		}

		if lastTemp == 0 {
			lastTemp = temp
			fmt.Printf("Initial CPU temperature: %.2f°C\n", temp)
			continue
		}

		if shouldNotify(temp, lastTemp) {
			if err := notifyUser(temp); err != nil {
				fmt.Printf("Error sending notification: %v\n", err)
				return
			}
		} else {
			fmt.Printf("CPU temperature is stable: %.2f°C\n", temp)
		}

		lastTemp = temp
		time.Sleep(500 * time.Millisecond) // Sleep for 250 milliseconds
	}
}
