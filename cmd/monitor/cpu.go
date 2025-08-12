package main

import (
	"context"
	"fmt"
	"log/slog"
	"sort"

	"github.com/jacobbrewer1/sensor-monitor/pkg/alerts"
	"github.com/jacobbrewer1/web"
	"github.com/jacobbrewer1/web/logging"
)

// watchCPUCoreTemperaturesTask wraps around the watchCPUCoreTemperatures function
// to provide it as an asynchronous task for the web application.
func (a *App) watchCPUCoreTemperaturesTask(l *slog.Logger) web.AsyncTaskFunc {
	return func(ctx context.Context) {
		watchCPUCoreTemperatures(
			ctx,
			l,
			a.alerter,
			a.tempChan,
			a.errChan,
		)
	}
}

// watchCPUCoreTemperatures listens for CPU core temperature updates and processes them.
func watchCPUCoreTemperatures(
	ctx context.Context,
	l *slog.Logger,
	alerter alerts.Alerter,
	tempChan <-chan map[int]float64,
	errChan <-chan error,
) {
	var lastTemps map[int]float64

	for {
		select {
		case temps := <-tempChan:
			// Process temperatures to check for drastic increases
			processCPUTemperatures(ctx, l, alerter, temps, lastTemps)
			lastTemps = temps

		case err := <-errChan:
			if err != nil {
				// Handle the error, e.g., log it or notify the user
				l.Error("Error receiving CPU core temperatures",
					logging.KeyError, err,
				)

				alerter.Alert(
					"Reading CPU Core Temperatures Failed",
					fmt.Sprintf("An error occurred while reading CPU core temperatures: %s", err),
				)
			}

		case <-ctx.Done():
			return
		}
	}
}

// processCPUTemperatures processes the CPU core temperatures and sends notifications if needed.
func processCPUTemperatures(
	ctx context.Context,
	l *slog.Logger,
	alerter alerts.Alerter,
	currentTemps map[int]float64,
	prevTemps map[int]float64,
) {
	// If no previous temperatures, nothing to compare against
	if prevTemps == nil || len(prevTemps) == 0 {
		return
	}

	spikingCores := make(map[int]string)

	// Check each CPU core for drastic temperature increases
	for coreID, currentTemp := range currentTemps {
		select {
		case <-ctx.Done():
			return
		default:
			prevTemp, exists := prevTemps[coreID]
			if !exists {
				// No previous temperature for this core, skip comparison
				continue
			}

			// Calculate percentage increase
			// Avoid division by zero
			if prevTemp <= 0 {
				continue
			}

			percentageIncrease := ((currentTemp - prevTemp) / prevTemp) * 100

			// Check if increase is greater than 5%
			if percentageIncrease > 15.0 {
				// Use tempWithBuffer to create a threshold with 2% buffer to avoid false alarms
				thresholdTemp := tempWithBuffer(prevTemp, 2)

				// Only alert if current temp exceeds the buffered threshold
				if currentTemp > thresholdTemp {
					l.Warn("Drastic CPU temperature increase detected",
						slog.Int("core_id", coreID),
						slog.Float64("previous_temp", prevTemp),
						slog.Float64("current_temp", currentTemp),
						slog.Float64("threshold_temp", thresholdTemp),
						slog.Float64("percentage_increase", percentageIncrease),
					)

					spikingCores[coreID] = fmt.Sprintf("%.2f -> %.2f", prevTemp, currentTemp)
				}
			}
		}
	}

	// If any cores are spiking, send an alert
	if len(spikingCores) > 0 {
		alertMessage := "Drastic CPU temperature increase detected on cores: "

		// Extract and sort core IDs
		coreIDs := make([]int, 0, len(spikingCores))
		for coreID := range spikingCores {
			coreIDs = append(coreIDs, coreID)
		}
		sort.Ints(coreIDs)

		// Build the alert message with sorted core IDs
		for _, coreID := range coreIDs {
			temp := spikingCores[coreID]
			alertMessage += fmt.Sprintf("Core %d: %s; ", coreID, temp)
		}

		alerter.Notify(
			"Drastic CPU Temperature Increase Detected",
			alertMessage,
		)
	}
}

// tempWithBuffer is a helper function that returns the temperature adjusted by a buffer percentage.
func tempWithBuffer(temp float64, bufferPercentage int) float64 {
	buffer := (temp * float64(bufferPercentage)) / 100.0
	return temp + buffer
}
