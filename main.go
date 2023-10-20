package main

import (
	"fmt"

	"github.com/sonnyb378/regexpchalleges/problems/easy"
	"github.com/sonnyb378/regexpchalleges/problems/hard"
	"github.com/sonnyb378/regexpchalleges/problems/intermediate"
)

func main() {

	// EASY
	// fmt.Println("Extract all words")
	// fmt.Println("1.", easy.ExtractAllWords("Hello World! This is a simple text."))

	// fmt.Println("Digits Only")
	// fmt.Println("2. a12345", easy.OnlyDigits("a12345"))
	// fmt.Println("2. 12345", easy.OnlyDigits("12345"))
	// fmt.Println("2. 123b45", easy.OnlyDigits("123b45"))

	// fmt.Println("Extract email addresses in a string")
	// fmt.Println("3.", easy.ExtractEmailAddress("Please contact support@example.com for assistance or info@domain.co.uk."))

	fmt.Println("Check if string is a valid url")
	fmt.Println("4. \"http://www.info.co.uk\":", easy.IsValidURL("http://www.info.co.uk"))
	fmt.Println("4. \"http://www.info.co.uk.\":", easy.IsValidURL("http://www.info.co.uk."))
	fmt.Println("========================")
	fmt.Println("")

	// fmt.Println("Extract words that are in uppercase")
	// fmt.Println("HELLO, World! THIS Is a TEXT.")
	// fmt.Println("5. ", easy.ExtractUppercaseWords("HELLO, World! THIS Is a TEXT."))

	// fmt.Println("Check if a string contains at least one uppercase letter")
	// fmt.Println("6. HelloWorld: ", easy.HasUppercase("HelloWorld"))
	// fmt.Println("6. helloworld: ", easy.HasUppercase("helloworld"))
	// fmt.Println("6. hellowOrld: ", easy.HasUppercase("hellowOrld"))

	// fmt.Println("Extract phone numbers in a string")
	// input1 := "Contact us at (123) 456-7890 or (987) 654-3210. For legal concerns, contact us at 123-456-7890, (123)4567890, or (123)-456-7890"
	// result1, err1 := easy.ExtractPhoneNumber(input1)
	// fmt.Println("7. Contact us at (123) 456-7890 or (987) 654-3210. : ", result1, err1)

	// fmt.Println("Check if a string contains given word (case insensitive)")
	// input1 := "I like applEs"
	// fmt.Println("8. ", easy.FindWord(input1, "apple"))
	// input2 := "But oranges are my favorite"
	// fmt.Println("8. ", easy.FindWord(input2, "apple"))

	// INTERMEDIATE
	// fmt.Println("Extract all HTML tags")
	// input := "<html><body><h1>Hello</h1><p>This is a <b>sample</b> text.</p></body></html>"
	// result, err := intermediate.ExtractAllHTMLTags(input)
	// fmt.Println("9. ", result, err)
	// input1 := "<Button>Custom Button</Button>"
	// result1, err1 := intermediate.ExtractAllHTMLTags(input1)
	// fmt.Println("9. ", result1, err1)
	// input2 := "this is just a string"
	// result2, err2 := intermediate.ExtractAllHTMLTags(input2)
	// fmt.Println("9. ", result2, err2)

	fmt.Println("Replace all occurrences of the word")
	input := "I have a cat. I love cats. Cats are the best!"
	result, err := intermediate.FindAndReplace(input, "cat", "dog")
	fmt.Println("Replace \"cat\" with \"dog\"")
	fmt.Println("Input: ", input)
	fmt.Println("10. ", result, err)
	fmt.Println("========================")
	fmt.Println("")

	// fmt.Println("Extract IPV4 addresses")
	// input := "192.168.1.1 is a valid IP address, but 300.400.500.600 is not."
	// result, err := intermediate.ExtractIPAddress(input)
	// fmt.Println("11. ", result, err)

	// fmt.Println("Check if a string is a valid hexadecimal color code")
	// result1, err1 := intermediate.IsValidColorCode("#FFA500")
	// fmt.Println("12. ", result1, err1)
	// result2, err2 := intermediate.IsValidColorCode("#XYZ123")
	// fmt.Println("12. ", result2, err2)
	// result3, err3 := intermediate.IsValidColorCode("#fff")
	// fmt.Println("12. ", result3, err3)

	// fmt.Println("Extract all sentences from a given text (sentences end with \".\", \"!\", or \"?\")")
	// result1, err1 := intermediate.ExtractSentences("This is a sentence. So is this! What is this?")
	// fmt.Println("13. ", result1, err1)
	// result2, err2 := intermediate.ExtractSentences("This is a sentence So is this What is this")
	// fmt.Println("13. ", result2, err2)

	// fmt.Println("Validate a string as a valid email address")
	// result1 := intermediate.IsValidEmail("support@example.com")
	// fmt.Println("14. ", result1)
	// result2 := intermediate.IsValidEmail("info@example.co.uk")
	// fmt.Println("14. ", result2)
	// result3 := intermediate.IsValidEmail("info@example.co.uk.")
	// fmt.Println("14. ", result3)
	// result4 := intermediate.IsValidEmail("info@example")
	// fmt.Println("14. ", result4)
	// result5 := intermediate.IsValidEmail("info_email")
	// fmt.Println("14. ", result5)
	// result6 := intermediate.IsValidEmail("user.name@google.com.ph")
	// fmt.Println("14. ", result6)
	// result7 := intermediate.IsValidEmail("user-name@google-com.ph")
	// fmt.Println("14. ", result7)

	// fmt.Println("Extract all words that contain at least one vowel (a, e, i, o, u).")
	// result1, err1 := intermediate.ExtractWordsWithVowels("This is a simple example.")
	// fmt.Println("15. ", result1, err1)
	// result2, err2 := intermediate.ExtractWordsWithVowels("ths s smpl xmpl ds nt hv vwl.")
	// fmt.Println("15. ", result2, err2)

	// fmt.Println("Check if a string is a valid date in the format \"YYYY-MM-DD\"")
	// result1, err1 := intermediate.IsValidDate("2023-10-06")
	// fmt.Println("16. ", result1, err1)
	// result2, err2 := intermediate.IsValidDate("06-10-2023")
	// fmt.Println("16. ", result2, err2)
	// result3, err3 := intermediate.IsValidDate("2023-01-09")
	// fmt.Println("16. ", result3, err3)
	// result4, err4 := intermediate.IsValidDate("2023-1-9")
	// fmt.Println("16. ", result4, err4)

	// HARD
	// fmt.Println("Extract all URLs from a given text, including those with query parameters and anchors.")
	// result1, err1 := hard.ExtractCompleteURL("Visit https://www.example.com/page?param=value#section or https://another.example/page2#top or https://www.example.com.")
	// fmt.Println("17.", result1, err1)

	// fmt.Println("Extract all valid ISBN-10 and ISBN-13 numbers from a given text (ISBN-10: 0-123456-78-9, ISBN-13: 978-0-123456-78-9).")
	// input := "Check out ISBN-10: 0-123456-78-9, ISBN-11: 978-0-123456-78-9 and ISBN-13: 978-0-123456-78-9."
	// result, err := hard.ExtractAllISBNNumbers(input)
	// fmt.Println("18.", result, err)

	// fmt.Println("Extract all HTML attributes and their values from a given HTML tag.")
	// input1 := "<a href='https://www.example.com' title='Example'>Link</a>"
	// result1, err1 := hard.ExtractHTMLAttributesAndValues(input1)
	// fmt.Println("19.", result1, err1)
	// input2 := "<ol><li class='first_child'>1</li><li class='selected'>2</li><li>3</li></ol>"
	// result2, err2 := hard.ExtractHTMLAttributesAndValues(input2)
	// fmt.Println("19.", result2, err2)
	// input3 := "<Button>Click Me!</Button>"
	// result3, err3 := hard.ExtractHTMLAttributesAndValues(input3)
	// fmt.Println("19.", result3, err3)

	// fmt.Println("Extract all SQL queries (SELECT, INSERT, UPDATE, DELETE) from a given SQL script.")
	// input1 := "SELECT * FROM users; INSERT INTO orders (product) VALUES ('Widget');"
	// result1, err1 := hard.ExtractSQL(input1)
	// fmt.Println("20.", result1, err1)

	// fmt.Println("Extract all valid phone numbers, including international formats (e.g., +1 (123) 456-7890), from a given text.")
	// input1 := "Contact us at +1 (123) 456-7890 or 555-123-4567."
	// result1, err1 := hard.ExtractValidPhoneNumbers(input1)
	// fmt.Println("21.", result1, err1)

	// fmt.Println("Check if a string is a valid XML document (contains proper opening and closing tags).")
	// input1 := "<root><element>Value</element></root>" // true: root=2, element=2
	// result1, err1 := hard.IsValidXMLDoc(input1)
	// fmt.Println("22A.", result1, err1)
	// input2 := "<root><element>Value</root>" // false: root=2, element=1
	// result2, err2 := hard.IsValidXMLDoc(input2)
	// fmt.Println("22B.", result2, err2)
	// input3 := "<root><element>Value</root></root>" // false: root=3, element=1
	// result3, err3 := hard.IsValidXMLDoc(input3)
	// fmt.Println("22C.", result3, err3)
	// input4 := "<root><element>Value</ele></root>" // false: root=2, element=1, ele=1
	// result4, err4 := hard.IsValidXMLDoc(input4)
	// fmt.Println("22C.", result4, err4)

	// fmt.Println("Extract all words that are palindromes (read the same forwards and backwards) from a given text.")
	// input1 := "level kayak radar text"
	// result1, err1 := hard.ExtractPalindromes(input1)
	// fmt.Println("23A.", result1, err1)
	// input2 := "levl kyak rada tex"
	// result2, err2 := hard.ExtractPalindromes(input2)
	// fmt.Println("23B.", result2, err2)

	fmt.Println("Replace all URLs in a given text with \"[Link Removed]\" while keeping the text intact.")
	input1 := "Visit https://www.example.com/page or http://another.example. https://google.com"
	result1, err1 := hard.RemoveLinks(input1)
	fmt.Println("24.", result1, err1)
	input2 := "Visit https://www.example./page or http://another.example."
	result2, err2 := hard.RemoveLinks(input2)
	fmt.Println("24.", result2, err2)
	input3 := "Visit this is not a link or this is not a link."
	result3, err3 := hard.RemoveLinks(input3)
	fmt.Println("24.", result3, err3)
	fmt.Println("========================")
	fmt.Println("")
}
