/*
Write a function that computes the running total of a list.
*/
package main

import "fmt"

func RunningTotal(list []int) {
	total := 0
	for _, n := range list {
		total += n
	}
	fmt.Print(total)
}

func main() {
	list := []int{
		1, 2, 3, 4, 5,
		6, 7, 8, 9, 10,
		11, 12, 13, 14,
	}

	RunningTotal(list)
}
