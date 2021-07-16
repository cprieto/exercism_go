package binarysearch

// SearchInts searchs an ordered slice of integers for elem using binary search
func SearchInts(items []int, elem int) int {
	var rec func(int, int) int

	rec = func(start, end int) int {
		delta := end - start

		// empty list, base case
		if delta == 0 {
			return -1
		}

		middle := delta/2 + start
		found := items[middle]

		switch {
		case found == elem:
			return middle
		case found < elem:
			return rec(middle+1, end)
		case found > elem:
			return rec(start, middle)
		default:
			return -1 // This should never happen anyway
		}
	}

	return rec(0, len(items))
}
