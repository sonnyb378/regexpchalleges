package hard

import (
	"regexp"
	"strings"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func ExtractAllISBNNumbers(input string) ([]string, error) {

	// Problem: Extract all valid ISBN-10 and ISBN-13 numbers from a given text (ISBN-10: 0-123456-78-9, ISBN-13: 978-0-123456-78-9).
	// Input: "Check out ISBN-10: 0-123456-78-9 and ISBN-13: 978-0-123456-78-9."
	// Output: ["0-123456-78-9", “978-0-123456-78-9"]

	isbn := []string{}
	re := regexp.MustCompile(`(?:ISBN-(?:10|13):)\s(?:-?\d+){1,}`)
	if re.MatchString(input) {
		matched := re.FindAllString(input, -1)
		for i := 0; i < len(matched); i++ {
			num := strings.Split(matched[i], " ")
			isbn = append(isbn, num[1])
		}
		return isbn, nil
	}

	return []string{}, &shared.CustomError{
		Message: "No valid ISBN-10 or ISBN-13 numbers found!",
	}
}
