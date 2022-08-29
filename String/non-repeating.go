package main

import "fmt"

// Example 1:
//
// Input: s = "leetcode"
// Output: 0
//
// Example 2:
//
// Input: s = "loveleetcode"
// Output: 2
//
// Example 3:
//
// Input: s = "aabb"
// Output: -1
func firstUniqChar(s string) int {
	count := make(map[rune]int)
	for _, v := range s {
		count[v]++
	}
	for i, v := range s {
		if count[v] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("aabb"))
}
