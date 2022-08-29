package main

import (
	"fmt"
	"sort"
)

func findDisappearedNumbers(nums []int) []int {
	sort.Ints(nums)
	var ans []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
