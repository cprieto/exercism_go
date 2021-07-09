package variablelengthquantity

func DecodeVarint(input []byte) ([]uint32, error) {
	result := []uint32{}
	var current uint32
	for _, val := range input {
		current |= uint32(val & 0x7f)
		if val&0x80 == 0 {
			result = append(result, current)
			current = 0
		}
		current <<= 7
	}

	return result, nil
}

func EncodeVarint(input []uint32) []byte {
	res := []byte{}
	for _, n := range input {
		var cont uint32
		var varint []byte
		for {
			varint = append([]byte{byte(cont | (n & 0x7F))}, varint...)
			n >>= 7
			if n == 0 {
				res = append(res, varint...)
				break
			}
			cont = 0x80
		}
	}
	return res
}
