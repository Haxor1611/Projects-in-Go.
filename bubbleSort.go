package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the array elements: ")
	elements := []int{1, 6, 3, 5, 4}
	sortedList := bubbleSort(elements)
	fmt.Println(sortedList)
}
func bubbleSort(list []int) []int {
	var count int = 0

	for j := 0; j <= len(list)-2; j++ {
		for i := 0; i <= len(list)-2; i++ {
			count = count + 1
			if list[i] > list[i+1] {
				var temp int
				temp = list[i+1]
				list[i+1] = list[i]
				list[i] = temp
			}
		}
	}
	return list
}
