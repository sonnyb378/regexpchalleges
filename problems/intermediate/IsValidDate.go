package intermediate

import (
	"regexp"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func IsValidDate(date string) (bool, error) {
	// Problem: Check if a string is a valid date in the format "YYYY-MM-DD".
	// Input (Correct): "2023-10-06"
	// Input (Incorrect): "06-10-2023"
	// Input (Incorrect): "2023-01-09"
	// Input (Incorrect): "2023-1-6"
	// Output (Correct): true
	// Output (Incorrect): false
	// Output (Incorrect): true
	// Output (Incorrect): false
	re := regexp.MustCompile(`[\d]{4}-(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2\d|3[0-1])`)
	if re.MatchString(date) {
		return true, nil
	}
	return false, &shared.CustomError{
		Message: "Invalid date format!",
	}
}
