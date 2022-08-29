package main

import "fmt"

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true
// Example 2:

// Input: s = "rat", t = "car"
// Output: false

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap := make(map[rune]int)

	for _, c := range s {
		hashMap[c]++
	}

	for _, c := range t {
		hashMap[c]--
	}

	for _, v := range s {
		if hashMap[v] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("aa", "bb"))
}
