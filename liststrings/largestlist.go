/*
Write a function that returns the largest element in a list.
*/
package main

import (
	"fmt"
)

func main() {
	list := []int{
		1, 2, 3, 4, 5,
		6, 33, 8, 9, 10,
		11, 12, 13, 14,
	}

	largest := 0
	for _, i := range list {
		if i > largest {
			largest = i
		}
	}
	fmt.Print(largest)
}
