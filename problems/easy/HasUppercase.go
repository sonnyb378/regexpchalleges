package easy

import (
	"regexp"
)

func HasUppercase(s string) bool {
	// Problem: Check if a string contains at least one uppercase letter.
	// Input (Correct): "HelloWorld"
	// Input (Incorrect): "helloworld"
	// Output (Correct): true
	// Output (Incorrect): false
	re := regexp.MustCompile(`[A-Z]+`)
	return re.MatchString(s)
}
