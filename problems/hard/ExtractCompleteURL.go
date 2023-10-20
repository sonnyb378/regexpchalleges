package hard

import (
	"regexp"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func ExtractCompleteURL(input string) ([]string, error) {

	// Problem: Extract all URLs from a given text, including those with query parameters and anchors.
	// Input: "Visit https://www.example.com/page?param=value#section or https://another.example/page2#top."
	// Output: ["https://www.example.com/page?param=value#section", "https://another.example/page2#top"]
	// Input: "https://www.example.com"
	// Output: [], Error("Invalid URL. No query parameters or anchors were provided.")

	re := regexp.MustCompile(`(?:http|https)\:\/{1,2}(?:[w]{3}|[\w\d-]+)(?:\.\w+){1,}/(?:(?:\?|\w+\?)|\w+)(?:\w+=\w+|\d+){1,}(?:(?:\&\w+=\w+)?){1,}(?:\#\w+){1,}`)

	if re.MatchString(input) {
		return re.FindAllString(input, -1), nil
	}

	return []string{}, &shared.CustomError{
		Message: "No valid URLs with query string and anchor found!",
	}
}
