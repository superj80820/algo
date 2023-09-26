// time complexity: O(n)
// space complexity: O(1)
// func moveZeroes(nums []int) {
// 	for i, j := 0, 0; i < len(nums) && j < len(nums); {
// 		if nums[i] != 0 {
// 			i++
// 			j = i
// 		} else {
// 			for j < len(nums) && nums[j] == 0 {
// 				j++
// 			}
// 			if j < len(nums) {
// 				nums[i], nums[j] = nums[j], nums[i]
// 				i++
// 			}
// 		}
// 	}
// }

// time complexity: O(n)
// space complexity: O(1)
func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}