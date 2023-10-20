package intermediate

import (
	"regexp"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func ExtractIPAddress(input string) ([]string, error) {
	// Problem: Extract all valid IPv4 addresses from a given text.
	// Input: "192.168.1.1 is a valid IP address, but 300.400.500.600 is not."
	// Output: [“192.168.1.1"]
	re := regexp.MustCompile(`(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)){3}`)
	if re.MatchString(input) {
		return re.FindAllString(input, -1), nil
	}

	return []string{}, &shared.CustomError{
		Message: "No valid ip addresses found!",
	}
}
