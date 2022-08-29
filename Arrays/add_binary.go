package main

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	var ans strings.Builder
	carry := 0
	reverse([]rune(a))
	reverse([]rune(b))
	maxlen := max(len(a), len(b))
	for i := 0; i < maxlen; i++ {
		var digitA, digitB int
		if i < len(a) {
			digitA = int(a[i] - '0')
		} else {
			digitA = 0
		}
		if i < len(b) {
			digitB = int(a[i] - '0')
		} else {
			digitB = 0
		}
		total := digitA + digitB + carry
		char := rune(total % 2)
		ans.WriteString(fmt.Sprintf("%d", char))
		carry = total / 2
	}
	var res string
	if carry == 1 {
		res = "1" + ans.String()
	}
	return res
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func reverse(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
