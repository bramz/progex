/*
Write a function that merges two sorted lists into a new sorted list. [1,4,6],[2,3,5] → [1,2,3,4,5,6].
You can do this quicker than concatenating them followed by a sort.
*/
package main

import (
	"fmt"
	"sort"
)

func ConcatMerge(list1 []int, list2 []int) []int {
	var list []int
	sort.Ints(list1)
	sort.Ints(list2)
	list = append(list1, list2...)
	sort.Ints(list)
	return list
}

func main() {
	list1 := []int{
		1, 3, 4, 4,
		5, 6, 7, 8,
	}

	list2 := []int{
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	fmt.Print(ConcatMerge(list1, list2))
}
