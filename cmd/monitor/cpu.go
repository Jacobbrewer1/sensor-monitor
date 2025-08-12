package main

import (
	"context"
	"fmt"
	"log/slog"

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
	for {
		select {
		case temps := <-tempChan:
			// Process the temperature data
			// For example, log it or send it to a web client
			_ = temps // Placeholder for actual processing logic

			fmt.Println("Received CPU core temperatures:", temps)

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
			return // Exit when the context is done
		}
	}
}
