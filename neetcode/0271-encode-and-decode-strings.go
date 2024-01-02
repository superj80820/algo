// tags: arrays&hashing

package main

import (
	"strconv"
	"strings"
)

type Codec struct{}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	res := make([]string, len(strs))
	for idx := range strs {
		res[idx] = strconv.Itoa(len(strs[idx])) + "#" + strs[idx]
	}
	return strings.Join(res, "")
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	var res []string
	for i := 0; i < len(strs); {
		j := i
		for strs[j] != '#' {
			j++
		}
		num, err := strconv.Atoi(strs[i:j])
		if err != nil {
			return []string{}
		}
		i = j + 1
		j = i + num
		res = append(res, strs[i:j])
		i = j
	}
	return res
}

// Decodes a single string to a list of strings.
// func (codec *Codec) Decode(strs string) []string {
// 	var res []string
// 	var curString string
// 	var curNumber int
// 	for idx := 0; idx < len(strs); idx++ {
// 		val := strs[idx]
// 		if val == '#' {
// 			for i := range curString {
// 				numIdx := len(curString) - 1 - i
// 				num, err := strconv.Atoi(string(curString[numIdx]))
// 				if err != nil {
// 					return []string{}
// 				}
// 				curNumber += num * int(math.Pow(10, float64(i)))
// 			}
// 			res = append(res, strs[idx+1:idx+1+curNumber])
// 			idx += curNumber
// 			curString = ""
// 			curNumber = 0
// 		} else {
// 			curString += string(val)
// 		}
// 	}
// 	return res
// }

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
