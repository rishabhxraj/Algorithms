package main

import (
	"fmt"
	"math"
)

func main() {
	len := lengthOfLongestSubstring("pwwkew")
	fmt.Println(len)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	i, j, max := 0, 0, 0
	set := make(map[rune]struct{})
	for i < len(s) {
		c := rune(s[i])
		for char := range set {
			if _, ok := set[char]; ok {
				delete(set, rune(s[j]))
				j++
			}
			set[c] = struct{}{}
			max = int(math.Max(float64(max), float64(i-j+1)))
			i++
		}
		return max
	}
	return 0
}
