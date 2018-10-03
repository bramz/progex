/*
Write a function that concatenates two lists. [a,b,c], [1,2,3] → [a,b,c,1,2,3]
*/
package main

import "fmt"

func ConcatLists(list1 []int, list2 []int) []int {
	/* not simplified
	var newlist []int
	for i := 0; i < len(list1); i++ {
		newlist = append(newlist, list1[i])
	}
	for n := 0; n < len(list2); n++ {
		newlist = append(newlist, list2[n])
	}
	return newlist
	*/
	list := append(list1, list2...)
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

	fmt.Print(ConcatLists(list1, list2))
}
