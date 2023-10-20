package hard

import (
	"regexp"
)

func ExtractSQL(input string) ([]string, error) {

	// Problem: Extract all SQL queries (SELECT, INSERT, UPDATE, DELETE) from a given SQL script.
	// Input: "SELECT * FROM users; INSERT INTO orders (product) VALUES ('Widget');"
	// Output: ["SELECT * FROM users;", "INSERT INTO orders (product) VALUES (‘Widget');"]

	re := regexp.MustCompile(`(?:SELECT|INSERT|UPDATE|DELETE)(?:[\w\d*()'"_\-\s]+);`)

	if re.MatchString(input) {
		return re.FindAllString(input, -1), nil
	}

	return []string{}, nil
}
