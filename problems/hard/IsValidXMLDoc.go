package hard

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sonnyb378/regexpchalleges/shared"
)

type CustomStruct struct {
	m map[string]int
}

func NewCustomStruct() *CustomStruct {
	return &CustomStruct{
		m: make(map[string]int),
	}
}

func (c *CustomStruct) Add(key string, value int) {
	c.m[key] = value
}
func (c *CustomStruct) Remove(key string) {
	delete(c.m, key)
}
func (c *CustomStruct) Get(key string) int {
	return c.m[key]
}
func (c *CustomStruct) Contains(key string) bool {
	fmt.Println("Contains: ", c.m[key])
	return false
}

func IsValidXMLDoc(input string) (bool, error) {

	// Problem: Check if a string is a valid XML document (contains proper opening and closing tags).
	// Input (Correct): "<root><element>Value</element></root>"
	// Input (Incorrect): "<root><element>Value</root>"
	// Output (Correct): true
	// Output (Incorrect): false

	newElement := NewCustomStruct()
	result := false

	errorMsg := "Error: Invalid XML document!"
	re := regexp.MustCompile(`<[\w\d]+>|</[\w\d]+>`)
	if re.MatchString(input) {
		matched := re.FindAllString(input, -1)
		if len(matched)%2 != 0 {
			return false, &shared.CustomError{
				Message: errorMsg,
			}
		}
		for _, v := range matched {
			replacer := strings.NewReplacer("<", "", ">", "", "/", "")
			el := replacer.Replace(v)
			newElement.Add(el, newElement.Get(el)+1)
		}
		if len(newElement.m)%2 != 0 {
			return false, &shared.CustomError{
				Message: errorMsg,
			}
		}
		for _, v := range newElement.m {
			if v%2 != 0 {
				return false, &shared.CustomError{
					Message: errorMsg,
				}
			}
		}

		result = true
	}

	return result, nil

}
