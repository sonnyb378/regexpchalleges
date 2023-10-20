package hard

import (
	"regexp"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func ExtractValidPhoneNumbers(input string) ([]string, error) {

	// Problem: Extract all valid phone numbers, including international formats (e.g., +1 (123) 456-7890), from a given text.
	// Input: "Contact us at +1 (123) 456-7890 or 555-123-4567."
	// Output: ["+1 (123) 456-7890", “555-123-4567"]

	re := regexp.MustCompile(`(?:\+\d\s\(\d{3}\)\s)??(?:\-?\d{3,4}){2,3}`)
	if re.MatchString(input) {
		return re.FindAllString(input, -1), nil
	}

	return []string{}, &shared.CustomError{
		Message: "No valid phone numbers found!",
	}
}
