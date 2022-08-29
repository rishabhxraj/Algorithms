package main

import "fmt"

// Example 1:

// Input: arr = [10,2,5,3]
// Output: true
// Explanation: N = 10 is the double of M = 5,that is, 10 = 2 * 5.
// Example 2:

// Input: arr = [7,1,14,11]
// Output: true
// Explanation: N = 14 is the double of M = 7,that is, 14 = 2 * 7.
// Example 3:

// Input: arr = [3,1,7,11]
// Output: false
// Explanation: In this case does not exist N and M, such that N = 2 * M.

func checkIfExist(arr []int) bool {
	exist := false
	for i := 0; i < len(arr); i++ {
		elem := 2 * arr[i]
		for _, a := range arr {
			if a == elem {
				exist = true
			}
		}
	}
	return exist
}

func main() {
	fmt.Println(checkIfExist([]int{-2, 0, 10, -19, 4, 6, -8}))
}
