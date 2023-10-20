package easy

import (
	"regexp"
)

func ExtractAllWords(input string) []string {
	// Problem: Extract all words (sequences of alphabetic characters) from a given text.
	// Input: "Hello World! This is a simple text."
	// Output: ["Hello", "World", "This", "is", "a", "simple", “text"]

	re := regexp.MustCompile(`[\w]+`)
	if re.MatchString(input) {
		return re.FindAllString(input, -1)
	}
	return []string{}
}
