package allergies

// Allergen is, well, as the name says, an allergen
type Allergen uint

func (a Allergen) isInScore(score uint) bool {
	return a&Allergen(score) != 0
}

const (
	// Eggs that thing birds and reptiles lays
	Eggs Allergen = 1 << iota
	// Peanuts that thing that you eat with bread
	Peanuts
	// Shellfish the delicious thing from the sea
	Shellfish
	// Strawberries is the red thing that we learn how to draw
	Strawberries
	// Tomatoes What makes pasta so yummy
	Tomatoes
	// Chocolate cacao stuff most people die for
	Chocolate
	// Pollen damn you Spring
	Pollen
	// Cats the saddest of all allergies
	Cats
)

var allergenToName = map[Allergen]string{
	Eggs:         "eggs",
	Peanuts:      "peanuts",
	Shellfish:    "shellfish",
	Strawberries: "strawberries",
	Tomatoes:     "tomatoes",
	Chocolate:    "chocolate",
	Pollen:       "pollen",
	Cats:         "cats",
}

var nameToAllergen = map[string]Allergen{
	"eggs":         Eggs,
	"peanuts":      Peanuts,
	"shellfish":    Shellfish,
	"strawberries": Strawberries,
	"tomatoes":     Tomatoes,
	"chocolate":    Chocolate,
	"pollen":       Pollen,
	"cats":         Cats,
}

var allAllergens = []Allergen{Eggs, Peanuts, Shellfish, Strawberries, Tomatoes, Chocolate, Pollen, Cats}

// Allergies returns the list of allergies for a given score
func Allergies(score uint) (result []string) {
	for _, v := range allAllergens {
		if v.isInScore(score) {
			result = append(result, allergenToName[v])
		}
	}
	return
}

// AllergicTo tells us if based in the score, we are allergic to that thing
func AllergicTo(score uint, substance string) bool {
	food, ok := nameToAllergen[substance]
	if !ok {
		return false
	}
	return food.isInScore(score)
}
