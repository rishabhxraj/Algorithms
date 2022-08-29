package main

import "fmt"

func replaceElements(arr []int) []int {
	var ans []int
	for i := 0; i < len(arr)-1; i++ {
		max := greatest(arr[i+1:])
		ans = append(ans, max)
	}
	ans = append(ans, -1)
	return ans
}

func greatest(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	res := replaceElements(arr)
	fmt.Println(res)
}
