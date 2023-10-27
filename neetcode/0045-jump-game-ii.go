// tags: greedy, dp(todo)

func jump(nums []int) int {
	var l, r, count int
	for r < len(nums)-1 {
		var nextR int
		for i := l; i <= r; i++ {
			nextR = max(nextR, i+nums[i])
			if nextR >= len(nums)-1 {
				break
			}
		}
		l, r = r+1, nextR
		count++
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func jump(nums []int) int {
//     if len(nums) == 1 {
//         return 0
//     }

//     queue := make([][2]int, 0, len(nums))

//     queue = append(queue, [2]int{0, nums[0]})
//     count := 1
//     for len(queue) != 0 {
//         curLen := len(queue)
//         var endIdx int
//         starIdx := queue[len(queue)-1][0]
//         for i := 0; i < curLen; i++ {
//             first := dequeue(&queue)
//             jumpIdx, jumpVal := first[0], first[1]
//             endIdx = max(endIdx, jumpIdx+jumpVal)
//             if endIdx >= len(nums)-1 {
//                 return count
//             }
//         }
//         for i := starIdx+1; i <= endIdx; i++ {
//             queue = append(queue, [2]int{i, nums[i]})
//         }
//         count++
//     }

//     return count
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

// func dequeue(queue *[][2]int) [2]int {
//     first := (*queue)[0]
//     *queue = (*queue)[1:]
//     return first
// }