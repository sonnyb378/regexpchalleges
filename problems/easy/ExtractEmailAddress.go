package easy

import (
	"regexp"
)

func ExtractEmailAddress(input string) []string {
	// Problem: Extract all email addresses from a given text.
	// Input: "Please contact support@example.com for assistance or info@domain.co.uk."
	// Output: ["support@example.com", "info@domain.co.uk"]

	email := []string{}

	re := regexp.MustCompile(`[\w_.+-]+@[\w-.]+\b`)
	if re.MatchString(input) {
		email = re.FindAllString(input, -1)
	}

	return email
}
