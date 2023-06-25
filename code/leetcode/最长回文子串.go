package leetcode

func longestPalindrome(s string) string {
	res := 0
	subStr := ""
	for i := range s {
		a, b := longest(s, i, i)
		if b-a > res {
			res = b - a
			subStr = s[a:b]
		}
		a, b = longest(s, i, i+1)
		if b-a > res {
			res = b - a
			subStr = s[a:b]
		}
		a, b = longest(s, i-1, i)
		if b-a > res {
			res = b - a
			subStr = s[a:b]
		}
	}
	return subStr
}
func longest(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}
		i--
		j++
	}
	return i + 1, j
}
