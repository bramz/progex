/*
Write a function that reverses a list, preferably in place
*/
package main

import (
	"fmt"
)

func ReverseList(list []int) []int {
	if len(list) == 0 {
		return list
	}
	return append(ReverseList(list[1:]), list[0])
}

func main() {
	list := []int{
		1, 2, 3, 4, 5,
		6, 33, 8, 9, 10,
		11, 12, 13, 14,
	}
	fmt.Print(ReverseList(list))
}
