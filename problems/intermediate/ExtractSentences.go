package intermediate

import (
	"regexp"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func ExtractSentences(input string) ([]string, error) {
	// Problem: Extract all sentences from a given text (sentences end with ".", "!", or "?").
	// Input: "This is a sentence. So is this! What is this?"
	// Output: ["This is a sentence.", "So is this!", "What is this?”]
	re := regexp.MustCompile(`[a-zA-Z0-9\s]+(?:[.!?])`)
	if re.MatchString(input) {
		return re.FindAllString(input, -1), nil
	}
	return []string{}, &shared.CustomError{
		Message: "Invalid input!",
	}
}
