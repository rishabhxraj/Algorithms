package main

func removeFirst(array []int, l int) {
	i := 2
	for i < len(array) {
		array[i-1] = array[i]
		i++
	}
	l--
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	l := len(array)
	removeFirst(array, l)
}
