// tags: math&geometry

// time complexity: (m*n)
// space complexity: (m+n)
// `m` is length of num1
// `n` is length of num2
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	productSlice := make([]byte, len(num1)+len(num2)+1)
	for i := range num1 {
		num1Digit := num1[len(num1)-1-i] - '0'
		for j := range num2 {
			num2Digit := num2[len(num2)-1-j] - '0'
			product := num1Digit * num2Digit
			productSlice[i+j] += product
			productSlice[i+j+1] += productSlice[i+j] / 10
			productSlice[i+j] = productSlice[i+j] % 10
		}
	}

	var rmZero int
	for i := len(productSlice) - 1; i >= 0; i-- {
		if productSlice[i] == 0 {
			rmZero++
		} else {
			break
		}
	}
	productSlice = productSlice[:len(productSlice)-rmZero]

	for i := range productSlice {
		productSlice[i] += '0'
	}

	for i, j := 0, len(productSlice)-1; i < j; i, j = i+1, j-1 {
		productSlice[i], productSlice[j] = productSlice[j], productSlice[i]
	}

	return string(productSlice)
}