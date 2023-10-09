package leetcode

import "bytes"

func addStrings(num1 string, num2 string) string {
	var buf bytes.Buffer
	m, n := len(num1)-1, len(num2)-1
	var a, b, c byte = 0, 0, 0
	var zero byte = '0'
	for n >= 0 || m >= 0 {
		a, b = 0, 0
		if m >= 0 {
			a = num1[m] - zero
			m--
		}
		if n >= 0 {
			b = num2[n] - zero
			n--
		}
		buf.WriteByte(zero + (a+b+c)%10)
		c = (a + b + c) / 10
	}
	if c > 0 {
		buf.WriteByte(zero + c)
	}
	res := buf.String()
	return Reverse(res)
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
