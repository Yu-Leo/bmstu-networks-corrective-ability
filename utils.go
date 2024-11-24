package main

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func powBinary(n uint64) uint64 {
	res := uint64(1)
	for i := uint64(1); i <= n; i++ {
		res <<= 1
	}
	return res
}

func GetDivisionRemainder(divisible, divider uint64) uint64 {
	difference := 1

	for difference >= 0 {
		difference = bitLen(divisible) - bitLen(divider)

		if difference > 0 {
			divisible = divisible ^ (divider << difference)
		} else {
			divisible = divisible ^ divider
		}

		difference = bitLen(divisible) - bitLen(divider)
	}

	return divisible
}

// bitLen - функция для получения количества битов в числе
func bitLen(n uint64) int {
	if n == 0 {
		return 0
	}
	bits := 0
	for n > 0 {
		bits++
		n >>= 1
	}
	return bits
}
