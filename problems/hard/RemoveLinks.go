package hard

import (
	"regexp"
	"strings"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func RemoveLinks(input string) (string, error) {

	// Problem: Replace all URLs in a given text with "[Link Removed]" while keeping the text intact.
	// Input: "Visit https://www.example.com/page or http://another.example."
	// Output: "Visit [Link Removed] or [Link Removed]."

	re := regexp.MustCompile(`(?i)(?:https?\:\/{2})(?:w{3})?\.?(?:\w+\.\w+)?(?:\/\w+)?`)
	var newStr string
	if re.MatchString(input) {
		replacedStr := re.ReplaceAllStringFunc(input, func(match string) string {
			newStr = match
			matchedArr := strings.Split(match, ".")
			if len(matchedArr) >= 2 && matchedArr[len(matchedArr)-1] != "" {
				newStr = "[Link Removed]"
			}
			return newStr
		})
		return replacedStr, nil
	}

	return "", &shared.CustomError{
		Message: "No links found in input string!",
	}
}
