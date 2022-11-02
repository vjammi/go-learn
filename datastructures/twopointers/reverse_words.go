package main

import (
	"fmt"
	"strings"
)

func reverseWords(sentence string) string {
	arr := strings.Split(sentence, " ")
	i := 0
	j := len(arr) - 1
	for i < j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	}
	return strings.Join(arr, " ")
}

func main() {
	arrOfStrs := []string{"Hello World", "Hello world golang", "The quick brown fox jumped over the lazy dog"}

	for i, str := range arrOfStrs {
		fmt.Println(i, " ", reverseWords(str))
	}
}
