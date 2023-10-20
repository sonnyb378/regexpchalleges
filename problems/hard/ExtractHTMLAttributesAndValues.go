package hard

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/sonnyb378/regexpchalleges/shared"
)

func ExtractHTMLAttributesAndValues(input string) (string, error) {
	// Problem: Extract all HTML attributes and their values from a given HTML tag.
	// Input: "<a href='https://www.example.com' title='Example'>Link</a>"
	// Output: [{"href": "https://www.example.com", "title": “Example"}]

	output := []map[string]string{}

	re := regexp.MustCompile(`(?:\w+)?(?:\=)(?:\'|\")[a-zA-Z0-9_./:-]+(?:\'|\")`)
	if re.MatchString(input) {
		matched := re.FindAllString(input, -1)
		for _, data := range matched {
			av := strings.Split(data, "=")
			output = append(output, map[string]string{
				av[0]: strings.Replace(av[1], "'", "", -1),
			})
		}
		result, err := json.Marshal(output)
		if err != nil {
			return "", &shared.CustomError{
				Message: err.Error(),
			}
		}

		return string(result), nil

	}

	return "", &shared.CustomError{
		Message: "No HTML attributes and values found!",
	}

}
