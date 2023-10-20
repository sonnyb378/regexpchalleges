package hard

import (
	"regexp"
	"strings"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func ExtractPalindromes(input string) ([]string, error) {

	// Problem: Extract all words that are palindromes (read the same forwards and backwards) from a given text.
	// Input: "level kayak radar text"
	// Output: ["level", "kayak", “radar"]

	palindromes := []string{}
	re := regexp.MustCompile(`[\w]+\b`)
	if re.MatchString(input) {
		matched := re.FindAllString(input, -1)
		for _, word := range matched {
			if IsPalindrome(word) {
				palindromes = append(palindromes, word)
			}
		}
	}
	if len(palindromes) > 0 {
		return palindromes, nil
	}
	return []string{}, &shared.CustomError{
		Message: "No palindromes found in given input string!",
	}
}

func IsPalindrome(word string) bool {
	var builder strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		builder.WriteString(string(word[i]))
	}
	return word == builder.String()
}
