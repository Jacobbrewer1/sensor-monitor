package alerts

import (
	"github.com/gen2brain/beeep"

	"github.com/jacobbrewer1/web/logging"
)

func (a *alerter) Notify(title, message string) {
	if err := beeep.Notify(
		title,
		message,
		"",
	); err != nil {
		a.l.Error("failed to send notification",
			logging.KeyError, err,
		)
		return
	}
}
