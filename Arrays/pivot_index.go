package main

import "fmt"

func pivotIndex(nums []int) int {
	leftSum := 0
	rightSum := sum(nums)

	for i := 0; i < len(nums); i++ {
		leftSum += nums[i]

		if leftSum == rightSum {
			return i
		}
		rightSum -= nums[i]
	}
	return -1
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}
