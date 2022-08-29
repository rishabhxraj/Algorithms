package main

import (
	"fmt"
	"math"
)

// Example 1:
//
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]
//
// Example 2:
//
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]
// Explanation: [9,4] is also accepted.
func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	freq1 := getFreq(nums1)
	freq2 := getFreq(nums2)

	for i := range freq1 {
		count := int(math.Min(float64(freq1[i]), float64(freq2[i])))
		for count > 0 {
			ans = append(ans, i)
			count--
		}
	}

	return ans
}

func getFreq(num []int) []int {
	freq := make([]int, 1001)
	for _, n := range num {
		freq[n]++
	}
	return freq
}

func main() {
	array1 := []int{4, 9, 5}
	array2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(array1, array2))
}
