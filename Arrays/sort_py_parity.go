package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	w := 0
	for r := 0; r < len(nums); r++ {
		if isEven(nums[r]) {
			nums[w], nums[r] = nums[r], nums[w]
			w++
		}
	}
	return nums
}

func isEven(num int) bool {
	return num%2 == 0
}

// Input: nums = [3,1,2,4]
// Output: [2,4,3,1]
// Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}
