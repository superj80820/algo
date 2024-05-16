// tags: 1d-dp, star3, easy, practice-count:3

// time complexity: O(n)
// space complexity: O(1)
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	n1, n2 := 1, 2
	for i := 2; i < n; i++ {
		temp := n2
		n2 = n1 + n2
		n1 = temp
	}

	return n2
}

// [8,5,3,2,1,1]
// time complexity: O(n)
// space complexity: O(n)
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n; i >= 0; i-- {
		if i+1 <= n {
			dp[i] += dp[i+1]
		}
		if i+2 <= n {
			dp[i] += dp[i+2]
		}
	}
	return dp[0]
}

// |  1  2
// 0  x  x  0
// 1  1  x  1
// 2  1  1  2
// 3  2  1  3
// 4  3  2  5
// 5  5  3  8
// time complexity: O(n)
// space complexity: O(n)
func climbStairs(n int) int {
	steps := []int{1, 2}
	dp := [][]*int{}
	for i := 0; i <= n; i++ {
		cur := make([]*int, len(steps)+1)
		for idx, step := range steps {
			if i-step < 0 {
				cur[idx] = nil
			} else if i-step == 0 {
				val := 1
				cur[idx] = &val
			} else {
				cur[idx] = dp[i-step][len(steps)]
			}
		}
		var sum int
		for idx := range steps {
			if cur[idx] != nil {
				sum += *cur[idx]
			}
		}
		cur[len(steps)] = &sum
		dp = append(dp, cur)
	}

	return *dp[n][len(steps)]
}
