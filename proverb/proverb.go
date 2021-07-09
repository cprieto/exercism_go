// Package proverb generate pieces of wisdom
package proverb

import "fmt"

// Proverb returns a proverb based in a list of nouns
func Proverb(rhyme []string) []string {
	items := len(rhyme)
	output := make([]string, 0, items)
	if items == 0 {
		return output
	}

	prev := rhyme[0]
	for idx := 1; idx < items; idx++ {
		current := rhyme[idx]
		message := fmt.Sprintf("For want of a %s the %s was lost.", prev, current)
		prev = current

		output = append(output, message)
	}
	output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return output
}
