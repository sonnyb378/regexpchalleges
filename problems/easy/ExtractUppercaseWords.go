package easy

import (
	"regexp"
)

func ExtractUppercaseWords(input string) []string {
	// Problem: Extract all words that are in uppercase from a given text.
	// Input: "HELLO, World! THIS Is a TEXT."
	// Output: ["HELLO", "THIS", “TEXT"]
	re := regexp.MustCompile(`[A-Z]+\b`)
	return re.FindAllString(input, -1)
}
