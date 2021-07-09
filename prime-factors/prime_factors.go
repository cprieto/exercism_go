package prime

// Factors returns the prime factors of number n
// taken from http://butunclebob.com/ArticleS.UncleBob.ThePrimeFactorsKata
func Factors(n int64) []int64 {
	result := make([]int64, 0)
	for i := int64(2); n > 1; i++ {
		for ; n%i == 0; n /= i {
			result = append(result, i)
		}
	}
	return result
}
