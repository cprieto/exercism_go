// railfence contains functions to work with rails fence cipher
package railfence

// Implementation O(n) based in paper by Godara, Kundu and Kaler
// see http://article.nadiapub.com/IJFGCN/vol11_no2/3.pdf

import "strings"

// Encode encodes a string using the rails fence cipher
func Encode(input string, rails int) string {
	entry := []rune(input)
	length := len(entry)

	var output strings.Builder

	for i := 1; i <= rails; i++ {
		flag := true
		pos := i

		for pos <= length {
			output.WriteRune(entry[pos-1])

			if i == 1 || i == rails {
				pos += (rails - 1) * 2
				continue
			}

			if flag {
				pos += (rails - i) * 2
			} else {
				pos += (i - 1) * 2
			}

			flag = !flag
		}
	}

	return output.String()
}

// Decode decodes a string using the rails fence cipher
func Decode(input string, rails int) string {
	entry := []rune(input)
	length := len(entry)

	output := make([]rune, length)

	cipher_char := 0
	for i := 1; i <= rails; i++ {
		pos := i
		flag := true

		for pos <= length {
			output[pos-1] = entry[cipher_char]
			cipher_char += 1

			if i == 1 || i == rails {
				pos += (rails - 1) * 2
				continue
			}

			if flag {
				pos += (rails - i) * 2
			} else {
				pos += (i - 1) * 2
			}

			flag = !flag
		}
	}

	return string(output)
}
