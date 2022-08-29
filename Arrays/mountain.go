package main

import "fmt"

// Example 1:

// Input: arr = [2,1]
// Output: false
// Example 2:

// Input: arr = [3,5,5]
// Output: false
// Example 3:

// Input: arr = [0,3,2,1]
// Output: true
func validMountainArray(arr []int) bool {
	i := 1
	for j := 1; j < len(arr); j++ {
		if arr[j] > arr[j-1] {
			i++
		} else {
			break
		}
	}

	if i == len(arr) || i == 1 {
		return false
	}

	for ; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
		} else {
			break
		}
	}
	return i == len(arr)
}

func main() {
	fmt.Println(validMountainArray([]int{4, 4, 3, 2, 1}))
}
