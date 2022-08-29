package main

import (
	"fmt"
	"math"
	"strconv"
)

// Example 1:
//
// Input: x = 123
// Output: 321
//
// Example 2:
//
// Input: x = -123
// Output: -321
//
// Example 3:
//
// Input: x = 120
// Output: 21
func reverse(x int) int {
	var isNeg bool
	if x < 0 {
		isNeg = true
	}
	absX := int(math.Abs(float64(x)))
	str := strconv.Itoa(absX)
	rev := reverseStr(str)
	revInt, _ := strconv.Atoi(rev)
	if isNeg {
		return -1 * revInt
	}
	return revInt
}

func reverseStr(s string) string {
	if len(s) <= 0 {
		return s
	}
	return reverseStr(string(s[1:])) + string(s[0])
}

func main() {
	fmt.Println(reverse(-123))
}
