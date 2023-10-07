package util

func FillNum(str string, fillLen int) string {
	strLen := len(str)
	if strLen >= fillLen {
		return str
	}
	var res string
	for i := 0; i < fillLen-strLen; i++ {
		res += "\u00A0"
	}
	return res + str
}
