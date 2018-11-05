/*
Write a function that rotates a list by k elements. For example [1,2,3,4,5,6] rotated by two becomes [3,4,5,6,1,2].
Try solving this without creating a copy of the list. How many swap or move operations do you need?
*/
package main

import "fmt"

func RotateList(list []int, k int) []int {
	rlen := len(list) - k%len(list)
	list = append(list[rlen:], list[:rlen]...)
	return list
}

func main() {
	list := []int{
		1, 2, 3, 4,
		5, 6, 7, 8,
	}
	fmt.Print(RotateList(list, 2), "\n")
	fmt.Print(RotateList(list, 3), "\n")
	fmt.Print(RotateList(list, 4), "\n")
}
