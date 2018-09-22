/*
Write a function that checks whether an element occurs in a list.
*/
package main

import "fmt"

func ElementInList(element int, list []int) bool {
	for _, elements := range list {
		if elements == element {
			return true
		}
	}
	return false
}

func main() {
	list := []int{
		1, 2, 3, 4, 5,
		6, 29, 83, 2,
		111, 1221, 312, 313,
	}
	element := 1221

	if ElementInList(element, list) == true {
		fmt.Printf("Element %d in list", element)
	} else {
		fmt.Printf("Element %d not in list", element)
	}
}
