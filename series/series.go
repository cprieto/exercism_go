package series

func All(n int, input string) []string {
	num := len(input)
	if n > num {
		return nil
	} else if n == num {
		return []string{input}
	}

	output := make([]string, 0)
	for idx := 0; idx <= num-n; idx++ {
		output = append(output, input[idx:idx+n])
	}
	return output
}

func UnsafeFirst(n int, input string) string {
	num := len(input)
	if n >= num {
		return input
	}

	return input[0:n]
}
