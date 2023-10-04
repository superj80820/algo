// First Factorial
// Have the function FirstFactorial(num) take the num parameter being passed and return the factorial of it. For example: if num = 4, then your program should return (4 * 3 * 2 * 1) = 24. For the test cases, the range will be between 1 and 18 and the input will always be an integer.
// Examples
// Input: 4
// Output: 24
// Input: 8
// Output: 40320
// Tags
// recursion math fundamentals free

package main

import "fmt"

func FirstFactorial(num int) int {
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	return res
}

import "fmt"

func FirstFactorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * FirstFactorial(num-1)
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(FirstFactorial(readline()))

}
