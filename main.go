package utf8len

var table = []int{
	// 0xxxx single byte
	1, 1, 1, 1,
	1, 1, 1, 1,
	1, 1, 1, 1,
	1, 1, 1, 1,
	// 10xxx continue byte (error)
	0, 0, 0, 0,
	0, 0, 0, 0,
	// 110xx two bytes
	2, 2, 2, 2,
	// 1110x three bytes
	3, 3,
	// 11110 four bytes
	4,
	// 11111 error
	0,
}

func FromFirstByte(c byte) int {
	return table[c>>3]
}
