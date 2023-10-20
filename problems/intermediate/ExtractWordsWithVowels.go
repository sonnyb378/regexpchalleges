package intermediate

import (
	"regexp"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func ExtractWordsWithVowels(input string) ([]string, error) {
	// Problem: Extract all words that contain at least one vowel (a, e, i, o, u).
	// Input: "This is a simple example."
	// Output: ["This", "is", "a", "simple", “example"]
	re := regexp.MustCompile(`[\w\d]*[aeiouAEIOU][\w\d]*`)
	if re.MatchString(input) {
		return re.FindAllString(input, -1), nil
	}

	return []string{}, &shared.CustomError{
		Message: "Input string does not contain words with vowels!",
	}
}
