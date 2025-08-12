package alerts

import (
	"github.com/gen2brain/beeep"

	"github.com/jacobbrewer1/web/logging"
)

// Alert sends an alert with the specified title and message.
func (a *alerter) Alert(title, message string) {
	if err := beeep.Alert(
		title,
		message,
		"",
	); err != nil {
		a.l.Error("failed to send alert",
			logging.KeyError, err,
		)
		return
	}
}
