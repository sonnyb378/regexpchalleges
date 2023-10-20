package easy

import (
	"regexp"
)

func OnlyDigits(n string) bool {
	// Problem: Check if a string contains only digits (0-9).
	// Input (Correct): "12345"
	// Input (Incorrect): "12A34"
	// Output (Correct): true
	// Output (Incorrect): false
	re := regexp.MustCompile(`(\b\d+\b)`)
	return re.MatchString(n)
}
