package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/gen2brain/beeep"

	"github.com/jacobbrewer1/sensor-monitor/pkg/alerts"
	"github.com/jacobbrewer1/sensor-monitor/pkg/sensors"
	"github.com/jacobbrewer1/sensor-monitor/pkg/utils"
	"github.com/jacobbrewer1/web"
	"github.com/jacobbrewer1/web/logging"
)

const (
	// appName is the name of the application.
	appName = "sensor-monitor"
)

type (
	// AppConfig holds the application configuration.
	AppConfig struct{}

	// App is the main application structure.
	App struct {
		base *web.App
		cfg  *AppConfig

		tempChan <-chan map[int]float64
		errChan  <-chan error
		alerter  alerts.Alerter
	}
)

// NewApp creates a new instance of App.
func NewApp(l *slog.Logger) (*App, error) {
	base, err := web.NewApp(l)
	if err != nil {
		return nil, fmt.Errorf("could not create new web app: %w", err)
	}

	cfg := new(AppConfig)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	return &App{
		base: base,
		cfg:  cfg,
	}, nil
}

// Start starts the application.
func (a *App) Start() error {
	if err := a.base.Start(
		web.WithMetricsEnabled(false),
		web.WithDependencyBootstrap(func(ctx context.Context) error {
			beeep.AppName = utils.PrettyName(appName)
			a.alerter = alerts.NewAlerter(
				logging.LoggerWithComponent(a.base.Logger(), "alerter"),
			)
			return nil
		}),
		web.WithDependencyBootstrap(func(ctx context.Context) error {
			a.tempChan, a.errChan = sensors.WatchCPUCoreTemperatureWithContext(ctx, time.Second)
			if a.tempChan == nil || a.errChan == nil {
				return errors.New("failed to initialize temperature monitoring channels")
			}
			return nil
		}),
		web.WithIndefiniteAsyncTask("watch-cpu-core-temperatures", a.watchCPUCoreTemperaturesTask(
			logging.LoggerWithComponent(a.base.Logger(), "watch-cpu-core-temperatures"),
		)),
	); err != nil {
		return fmt.Errorf("failed to start web app: %w", err)
	}
	return nil
}

// Shutdown gracefully shuts down the application.
func (a *App) Shutdown() {
	a.base.Shutdown()
}

// WaitForEnd blocks until the application is stopped.
func (a *App) WaitForEnd() {
	a.base.WaitForEnd(a.Shutdown)
}

func main() {
	l := logging.NewLogger(logging.WithAppName(appName))
	app, err := NewApp(l)
	if err != nil {
		l.Error("failed to create app",
			logging.KeyError, err,
		)
		panic(err)
	}

	if err := app.Start(); err != nil {
		l.Error("failed to start app",
			logging.KeyError, err,
		)
		panic(err)
	}

	app.WaitForEnd()
}
