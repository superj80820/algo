// tags: bit-manipulation

// time complexity: O(1)
// space complexity: O(1)
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		var bitMask uint32 = 1 << i
		if num&bitMask != 0 {
			res += 1 << (31 - i)
		}
	}
	return res
}