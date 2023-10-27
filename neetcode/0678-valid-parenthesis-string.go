// tags: greedy

func checkValidString(s string) bool {
	var minLeft, maxLeft int

	for _, val := range s {
		if val == '(' {
			minLeft++
		} else {
			minLeft--
		}
		if val == ')' {
			maxLeft--
		} else {
			maxLeft++
		}
		if minLeft < 0 {
			minLeft = 0
		}
		if maxLeft < 0 {
			return false
		}
	}
	return minLeft == 0 && maxLeft >= 0
}