// Package leap is a package for the Leap exercise in Exercism
package leap

// IsLeapYear tell us if a given year is a leap year or not
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
