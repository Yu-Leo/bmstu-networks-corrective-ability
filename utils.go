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

func getBinaryLength(digit uint64) uint64 {
	bitsNum := uint64(0)
	for ; digit/2 != 0; digit /= 2 {
		bitsNum++
	}
	bitsNum++
	return bitsNum
}

func IntToBytes(digit uint64) []byte {
	var res []byte
	for i := powBinary(getBinaryLength(digit) - 1); i > 0; i /= 2 {
		res = append(res, byte(digit/i))
		digit %= i
	}
	return res
}

func OperationO(a, b uint64) (uint64, uint64) {
	if a < b {
		return 0, a
	}

	var integer uint64
	aBytes := IntToBytes(a)
	bLen := getBinaryLength(b)
	var cur uint64

	aBytesPos := uint64(0)
	for ; aBytesPos < bLen; aBytesPos++ {
		cur <<= 1
		cur += uint64(aBytes[aBytesPos])
	}

	for ; aBytesPos <= uint64(len(aBytes)); aBytesPos++ {
		firstBitInCur := cur / powBinary(bLen-1)
		integer <<= 1
		integer += firstBitInCur

		if firstBitInCur == 1 {
			cur ^= b
		}
		if aBytesPos == uint64(len(aBytes)) {
			break
		}

		cur <<= 1
		cur += uint64(aBytes[aBytesPos])
	}

	return integer, cur
}
