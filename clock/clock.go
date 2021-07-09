package clock

import (
	"fmt"
	"math"
)

func mod(n, m int) int {
	return ((n % m) + m) % m  // This is a real mod, not a remainder!
}

type Clock struct {
	h int
	m int
}

func New(hh, mm int) Clock {
	hfm := math.Floor(float64(mm) / 60.0)  // We need the floor division between mm and 60
	hh = mod(hh + int(hfm), 24)
	mm = mod(mm, 60)
	return Clock{hh, mm}
}

func (clock *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.h, clock.m)
}

func (clock Clock) Add(mm int) Clock {
	mm = mm + clock.m
	hfm := math.Floor(float64(mm) / 60.0)  // This is hours from minutes

	clock.h = mod(clock.h + int(hfm), 24)
	clock.m = mod(mm, 60)

	return clock
}


func (clock Clock) Subtract(mm int) Clock {
	mm = clock.m - mm
	hfm := math.Floor(float64(mm) / 60.0)

	clock.h = mod(clock.h + int(hfm), 24)
	clock.m = mod(mm, 60)

	return clock
}
