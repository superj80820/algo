package main

import (
	"math"
)

// import "math"

// time complexity: O(n)
// space complexity: O(n)
// `n` is all num of sum two square number
func isHappy(n int) bool {
	check := make(map[int]bool)
	for {
		var curNum int
		for n != 0 {
			digit := n % 10
			n = n / 10
			curNum += int(math.Pow(float64(digit), 2))
		}
		if curNum == 1 {
			return true
		}
		if _, ok := check[curNum]; ok {
			return false
		}
		n = curNum
		check[n] = true
	}
}

// time complexity: O(n)
// space complexity: O(1)
// `n` is all num of sum two square number
func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = getNextNum(slow)
		fast = getNextNum(getNextNum(fast))
		if fast == 1 || slow == 1 {
			return true
		}
		if slow == fast {
			break
		}
	}
	return false
}

func getNextNum(n int) int {
	var curNum int
	for n != 0 {
		digit := n % 10
		n = n / 10
		curNum += int(math.Pow(float64(digit), 2))
	}
	return curNum
}
