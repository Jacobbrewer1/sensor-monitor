package alerts

import "log/slog"

type Alerter interface {
	// Alert sends an alert with the specified title and message.
	Alert(title, message string)

	// Notify sends a notification with the specified title and message.
	Notify(title, message string)
}

var _ Alerter = (*alerter)(nil)

type alerter struct {
	l *slog.Logger
}

// NewAlerter creates a new Alerter instance.
func NewAlerter(l *slog.Logger) Alerter {
	return &alerter{
		l: l,
	}
}
