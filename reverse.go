package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string
	fmt.Scanln(&input)
	fmt.Println(reverse(input))
}

func reverse(word string) string {
	size := len(word)
	charArray := []byte(word)
	reversedWord := strings.Builder{}
	for i := size - 1; i >= 0; i-- {
		reversedWord.WriteByte(charArray[i])
	}
	result := reversedWord.String()
	return result
}
