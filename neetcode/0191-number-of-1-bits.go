// tags: bit-manipulation, easy

// time complexity: O(1)
// space complexity: O(1)
func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		count += num % 2
		num = num >> 1
	}
	return count
}

// time complexity: O(1)
// space complexity: O(1)
func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

// time complexity: O(1)
// space complexity: O(1)
func hammingWeight(num uint32) int {
	var count int
	for i := 0; i < 32; i++ {
		var bitMask uint32
		bitMask = 1
		bitMask = bitMask << i
		if num&bitMask != 0 {
			count++
		}
	}
	return count
}