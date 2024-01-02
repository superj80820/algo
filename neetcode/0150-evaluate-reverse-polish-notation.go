// tags: stack, star1, medium

import "strconv"

// time complexity: O(n)
// space complexity: O(n)
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		if token == "+" {
			pop1, pop2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, pop2+pop1)
		} else if token == "-" {
			pop1, pop2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, pop2-pop1)
		} else if token == "*" {
			pop1, pop2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, pop2*pop1)
		} else if token == "/" {
			pop1, pop2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, pop2/pop1)
		} else {
			tokenInt, _ := strconv.Atoi(token)
			stack = append(stack, tokenInt)
		}
	}
	return stack[0]
}