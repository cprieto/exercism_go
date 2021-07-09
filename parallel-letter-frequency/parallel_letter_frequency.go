package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency returns the letter frequency of text in a concurrent way
// it returns data as a FreqMap
func ConcurrentFrequency(phrases []string) FreqMap {
	freqStream := make(chan map[rune]int, len(phrases))

	for _, entry := range phrases {
		go func(data string) {
			freqStream <- Frequency(data)
		}(entry)
	}

	result := make(FreqMap)
	for range phrases {
		for key, freq := range <- freqStream {
			result[key] += freq
		}
	}
	return result
}
