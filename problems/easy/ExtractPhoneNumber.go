package easy

import (
	"fmt"
	"regexp"
)

type PhoneNumberError struct {
	input string
}

func (e *PhoneNumberError) Error() string {
	return fmt.Sprintf("%v", e.input)
}

func ExtractPhoneNumber(input string) ([]string, error) {
	// Problem: Extract all phone numbers in the format (XXX) XXX-XXXX from a given text.
	// Input: "Contact us at (123) 456-7890 or (987) 654-3210."
	// Output: ["(123) 456-7890", "(987) 654-3210”]
	re := regexp.MustCompile(`\(\d{3}\)\s\d{3}-\d{4}`)
	if re.MatchString(input) {
		return re.FindAllString(input, -1), nil
	}
	return []string{}, &PhoneNumberError{
		input: "has no valid phone numbers",
	}
}
