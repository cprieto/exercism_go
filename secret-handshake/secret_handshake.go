package secret

import "math"

var code = map[uint]string{
	8: "jump",
	4: "close your eyes",
	2: "double blink",
	1: "wink",
}

func reverse(input []string) []string {
	length := len(input)
	if length == 0 {
		return input
	}

	for i := 0; i < length/2; i++ {
		j := length - i - 1
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func converter(input uint, current []string) []string {
	if input == 0 {
		return current
	}

	key := uint(1 << uint(math.Log2(float64(input))))
	current = append(current, code[key])
	next := input - key
	return converter(next, current)
}

func Handshake(input uint) []string {
	if input >= 16 {
		return converter(input-16, make([]string, 0))
	}
	return reverse(converter(input, make([]string, 0)))
}
