package main

import (
	"fmt"
)

// Example 1:
//
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// Example 2:
//
// Input: nums = [0]
// Output: [0]
func moveZeros(arr []int) []int {
	i := 0
	for j := 0; j < len(arr); j++ {
		if arr[j] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	return arr
}

func main() {
	array := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeros(array))
}
