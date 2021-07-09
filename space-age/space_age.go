package space

import "fmt"

// Planet represents a planet in the Solar System
type Planet string

const earthSeconds = 31557600

var planets = map[Planet]float64{
	Planet("Mercury"): 0.2408467,
	Planet("Venus"):   0.61519726,
	Planet("Earth"):   1.0,
	Planet("Mars"):    1.8808158,
	Planet("Jupiter"): 11.862615,
	Planet("Saturn"):  29.447498,
	Planet("Uranus"):  84.016846,
	Planet("Neptune"): 164.79132,
}

// Age returns the number of "ages" a number of seconds
// will be in a given planet in the Solar System
func Age(seconds float64, planet Planet) float64 {
	earthYears := seconds / earthSeconds
	if eq, ok := planets[planet]; ok {
		return earthYears / eq
	}
	panic(fmt.Sprintf("We don't know what planet is %s", planet))
}
