package intermediate

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func FindAndReplace(input string, find string, replaceWith string) (string, error) {
	// Problem: Replace all occurrences of the word "apple" with "banana" (case insensitive) in a given text.
	// Input: "I have an Apple. She has an apple. They like apples."
	// Output: "I have an Banana. She has an banana. They like bananas."
	re := regexp.MustCompile(`(?i)` + find)
	if re.MatchString(input) {

		replacedStr := re.ReplaceAllStringFunc(input, func(match string) string {
			newStr := replaceWith

			// get the first character of 'match'(find)
			firstChar := rune(match[0])

			// check if the first character is capitalized
			if unicode.IsUpper(firstChar) {

				// capitalize the first character of 'replaceWith' then get the rest of the substring
				newStr = strings.ToUpper(string(replaceWith[0])) + replaceWith[1:]
			}

			// return the replaced word
			return newStr
		})

		return replacedStr, nil
	}

	return "", &shared.CustomError{
		Message: fmt.Sprintf("\"%v\" not found in input string!", find),
	}
}
