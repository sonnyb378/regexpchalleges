package easy

import (
	"regexp"
)

func FindWord(input string, word string) bool {
	// Problem: Check if a string contains the the given word (apple) (case insensitive).
	// Input (Correct): "I like apples."
	// Input (Incorrect): "Oranges are my favorite."
	// Output (Correct): true
	// Output (Incorrect): false
	re := regexp.MustCompile(`(?i)` + word)
	return re.MatchString(input)

}
