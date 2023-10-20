package intermediate

import (
	"regexp"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func IsValidColorCode(input string) (bool, error) {
	// Problem: Check if a string is a valid hexadecimal color code (e.g., #RRGGBB or #RGB).
	// Input (Correct): "#FFA500"
	// Input (Incorrect): "#XYZ123"
	// Output (Correct): true
	// Output (Incorrect): false

	re := regexp.MustCompile(`#(?:[0-9a-fA-F]{1,2}){3}\b`)
	if re.MatchString(input) {
		return re.MatchString(input), nil
	}
	return re.MatchString(input), &shared.CustomError{
		Message: "input is not a valid hexadecimal color code!",
	}
}
