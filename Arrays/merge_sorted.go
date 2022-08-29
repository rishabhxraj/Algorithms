package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	for i < n {
		for j := 0; j < m+n; j++ {
			if nums1[j] <= nums2[i] {
				for k := len(nums1) - 1; k > j; k-- {
					nums1[j] = nums1[j-1]
				}
			}
		}
		i++
	}
}

//Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//Output: [1,2,2,3,5,6]
func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	m, n := 3, 3
	merge(arr1, m, arr2, n)
	fmt.Println(arr1)
}
