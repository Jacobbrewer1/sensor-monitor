package main

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

// alertUser sends a system notification to the user with the specified title and message.
func alertUser(title, message string) error {
	if err := beeep.Alert(
		title,
		message,
		"",
	); err != nil {
		return fmt.Errorf("failed to send alert: %w", err)
	}
	return nil
}

// notifyUser sends a notification to the user with the specified title and message.
func notifyUser(title, message string) error {
	if err := beeep.Notify(
		title,
		message,
		"",
	); err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}
	return nil
}
