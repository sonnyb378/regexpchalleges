package easy

import (
	"regexp"
	"unicode/utf8"
)

func IsValidURL(url string) bool {

	// Problem: Check if a string is a valid URL (starting with "http://" or "https://").
	// Input (Correct): "https://www.example.com"
	// Input (Incorrect): "ftp://example.com"
	// Output (Correct): true
	// Output (Incorrect): false

	re := regexp.MustCompile(`(http|https):\/\/[A-Za-z0-9.-]+\b`)
	if re.MatchString(url) {
		// input url length must be equal to the length of the matched string
		if utf8.RuneCountInString(url) == utf8.RuneCountInString(re.FindString(url)) {
			return true
		}
	}
	return false
}
