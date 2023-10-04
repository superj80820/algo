// Longest Word
// Have the function LongestWord(sen) take the sen parameter being passed and return the longest word in the string. If there are two or more words that are the same length, return the first word from the string with that length. Ignore punctuation and assume sen will not be empty. Words may also contain numbers, for example "Hello world123 567"
// Examples
// Input: "fun&!! time"
// Output: time
// Input: "I love dogs"
// Output: love
// Tags
// string manipulation searching free

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func LongestWord(sen string) string {
	var (
		res      string
		maxCount int
	)
	for _, word := range strings.Split(sen, " ") {
		var count int
		for _, char := range word {
			if unicode.IsLetter(char) || unicode.IsNumber(char) {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			res = word
		}
	}
	return res

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(LongestWord(readline()))

}
