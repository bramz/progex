/*
Write a function that combines two lists by alternatingly taking elements, e.g. [a,b,c], [1,2,3] → [a,1,b,2,c,3].
*/
package main

import "fmt"

func ConcatAlternate(list1 []int, list2 []int) []int {
	var list []int
	list = []int{
		list1[0], list2[0],
		list1[1], list2[1],
		list1[2], list2[2],
		list1[3], list2[3],
		list1[4], list2[4],
		list1[5], list2[5],
		list1[6], list2[6],
		list1[7], list2[7],
	}
	/* attempted with a loop, had issues iterating elements
	 * will return to this later, for now hand picking elements
	for i := 0; i < len(list1); i++ {
		list = []int{
			list1[i], list2[i],
		}
	}
	*/
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
	fmt.Print(ConcatAlternate(list1, list2))
}
