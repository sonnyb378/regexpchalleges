package intermediate

import (
	"regexp"
	"unicode/utf8"
)

func IsValidEmail(email string) bool {
	// Problem: Validate a string as a valid email address.
	// Input (Correct): "user@example.com"
	// Input (Incorrect): "invalid_email"
	// Output (Correct): true
	// Output (Incorrect): false

	re := regexp.MustCompile(`(?:[\w\d.-]+)@(?:[\w\d-]+(?:\.(?:[\w\d-]+\b)){1,})`)
	if re.MatchString(email) {
		if utf8.RuneCountInString(email) == utf8.RuneCountInString(re.FindString(email)) {
			return true
		}
	}

	return false
}
