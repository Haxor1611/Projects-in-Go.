package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{1, 3, 60, 140, 400}

	fmt.Println("Search: ")
	var x int
	fmt.Scanln(&x)
	result := search(arr, x)
	if result == -1 {
		fmt.Println("Element not present.")
	} else {
		s := strconv.Itoa(result + 1)
		fmt.Println("Element present at index: " + s)
	}
}
func search(arr []int, x int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == x {
			return i
		}
	}
	return -1
}
