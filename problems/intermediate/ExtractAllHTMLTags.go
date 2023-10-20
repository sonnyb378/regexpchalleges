package intermediate

import (
	"regexp"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func ExtractAllHTMLTags(input string) ([]string, error) {
	// Problem: Extract all HTML tags (e.g., <div>, <p>) from an HTML document.
	// Input: "<html><body><h1>Hello</h1><p>This is a <b>sample</b> text.</p></body></html>"
	// Output: ["<html>", "<body>", "<h1>", "</h1>", "<p>", "<b>", "</b>", "</p>", "</body>", “</html>"]

	re := regexp.MustCompile(`<[a-zA-Z0-9]+>|</[a-zA-Z0-9]+>`)
	if re.MatchString(input) {
		return re.FindAllString(input, -1), nil
	}

	return []string{}, &shared.CustomError{
		Message: "No valid HTML tags found!",
	}
}
