package main

func twoSum(numbers []int, target int) []int {
	leftIdx, rightIdx := 0, len(numbers)-1
	for leftIdx < rightIdx {
		sumNumber := numbers[leftIdx] + numbers[rightIdx]
		if sumNumber == target {
			return []int{leftIdx + 1, rightIdx + 1}
		} else if sumNumber < target {
			leftIdx++
		} else if sumNumber > target {
			rightIdx--
		}
	}
	return nil
}
