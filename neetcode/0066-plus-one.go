// tags: math&geometry

// time complexity: O(n)
// space complexity: O(n)
func plusOne(digits []int) []int {
	var carry int
	digits[len(digits)-1] += 1
	for i := len(digits) - 1; i >= 0; i-- {
		curSum := carry + digits[i]
		carry = curSum / 10
		digits[i] = curSum % 10
	}
	if carry == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

// time complexity: O(n)
// space complexity: O(1)
func plusOne(digits []int) []int {
	var carry int
	digits[len(digits)-1] += 1
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	for i := 0; i < len(digits); i++ {
		curSum := carry + digits[i]
		carry = curSum / 10
		digits[i] = curSum % 10
	}
	if carry == 1 {
		digits = append(digits, 1)
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}
