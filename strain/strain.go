// Package strain contains functions for handling strains
package strain

// Ints represents a sequence of integers
type Ints []int

// Strings represents a sequence of strings
type Strings []string

// Lists represents a sequence of integer lists
type Lists [][]int

// Keep filter a sequence of integers by those
// which op applies
func (l Ints) Keep(op func(int) bool) Ints {
	upTo := len(l)
	if upTo == 0 {
		return Ints(nil)
	}

	result := make([]int, 0)

	for n := 0; n < upTo; n++ {
		if e := l[n]; op(e) {
			result = append(result, e)
		}
	}

	return result
}

// Keep filter a sequence of strings by those
// which op applies
func (l Strings) Keep(op func(string) bool) Strings {
	upTo := len(l)
	if upTo == 0 {
		return Strings(nil)
	}

	result := make([]string, 0)

	for n := 0; n < upTo; n++ {
		if e := l[n]; op(e) {
			result = append(result, e)
		}
	}

	return result
}

// Keep filter a sequence of list of integers
// by those which op applies
func (l Lists) Keep(op func([]int) bool) Lists {
	upTo := len(l)
	if upTo == 0 {
		return Lists(nil)
	}
	result := make([][]int, 0)

	for n := 0; n < upTo; n++ {
		if e := l[n]; op(e) {
			result = append(result, e)
		}
	}

	return result
}

// Discard filter a sequence of integers by those
// which op does NOT apply
func (l Ints) Discard(op func(int) bool) Ints {
	return l.Keep(func(i int) bool { return !op(i) })
}

// Discard filter a squence of strings by those
// which op does NOT apply
func (l Strings) Discard(op func(string) bool) []string {
	return l.Keep(func(i string) bool { return !op(i) })
}

// Discard filter a squence of list of integers by those
// which op does NOT apply
func (l Lists) Discard(op func([]int) bool) [][]int {
	return l.Keep(func(i []int) bool { return !op(i) })
}
