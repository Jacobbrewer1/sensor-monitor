package utils

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// PrettyName converts a given name string into a more readable format for users
func PrettyName(name string) string {
	if name == "" {
		return "Unknown"
	}

	// Trim leading and trailing whitespace
	name = strings.TrimSpace(name)

	// Convert camelCase to snake_case by inserting underscores before uppercase letters
	// that are preceded by lowercase letters
	camelCaseRegex := regexp.MustCompile(`([a-z])([A-Z])`)
	name = camelCaseRegex.ReplaceAllString(name, "${1}_${2}")

	// Replace underscores with spaces and capitalize the first letter of each word
	name = strings.ReplaceAll(name, "_", " ")
	name = cases.Title(language.English).String(name)

	return name
}
