package isbn

import "regexp"

var validIsbn *regexp.Regexp = regexp.MustCompile(`(?mi)^[0-9 -X]+$`)
var cleanIsbn *regexp.Regexp = regexp.MustCompile(`(?m)[ -]`)

func IsValidISBN(isbn string) bool {
	if !validIsbn.MatchString(isbn) {
		return false
	}

	isbn = cleanIsbn.ReplaceAllString(isbn, "")
	if len(isbn) != 10 {
		return false
	}

	total := 0
	for idx, char := range isbn {
		factor := 10 - idx
		if char == 'x' || char == 'X' && factor != 1 {
			return false
		} else if char == 'x' || char == 'X' {
			total += factor * 10
		} else {
			total += factor * int(char-'0')
		}
	}
	return total%11 == 0
}
