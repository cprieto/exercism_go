// Package twofer is my answer for Exercism lesson 02
package twofer

import "fmt"

// ShareWith return the solution for Two-Fer question
// it will return the typical One for `name`, one for me
// if name is empty it will use `you` instead
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
