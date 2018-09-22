/*
Write a function that reverses a list, preferably in place
*/
package main

import (
	"fmt"
)

func ReverseList(list []int) []int {
	/* not in place
	if len(list) == 0 {
		return list
	}
	return append(ReverseList(list[1:]), list[0])
	*/
	for x, y := 0, len(list)-1; x < y; x, y = x+1, y-1 {
		list[x], list[y] = list[y], list[x]
	}
	return list
}

func main() {
	list := []int{
		1, 2, 3, 4, 5,
		6, 33, 8, 9, 10,
		11, 12, 13, 14,
	}
	fmt.Print(ReverseList(list))
}
