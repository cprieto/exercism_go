// Package sublist has functions to deal with sublists and its types
package sublist

// Relation is a type of relation between two lists
type Relation string

var (
	// Equal both lists are the same
	Equal = Relation("equal")
	// Sub list is a sublist
	Sub = Relation("sublist")
	// Super list is a superlist
	Super = Relation("superlist")
	// Unequal both list are unequal
	Unequal = Relation("unequal")
)

type checkFunction func(a, b []int) bool

func checkEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if b[i] != e {
			return false
		}
	}
	return true
}

func checkSublist(a, b []int) bool {
	if len(a) >= len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	la := len(a)
	lb := len(b)
	for idx := 0; idx <= lb-la; idx++ {
		piece := b[idx : idx+la]
		if checkEqual(a, piece) {
			return true
		}
	}
	return false
}

func checkSuperlist(a, b []int) bool {
	return checkSublist(b, a)
}

var checks = map[Relation]checkFunction{
	Equal: checkEqual,
	Sub:   checkSublist,
	Super: checkSuperlist,
}

// Sublist tells you what type of relation exist between two given lists
func Sublist(a, b []int) Relation {
	for knd, fn := range checks {
		if fn(a, b) {
			return knd
		}
	}

	return Unequal
}
