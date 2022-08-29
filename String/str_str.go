package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	for i, j := 0, l2-1; j < l1; {
		if needle == haystack[i:j+1] {
			return i
		}
		i++
		j++
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
}
