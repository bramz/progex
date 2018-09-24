/*
Write three functions that compute the sum of the numbers in a list: using a for-loop,
a while-loop and recursion. (Subject to availability of these constructs in your language of choice.)
Side note:
golangs for is while loop, so to expand a bit i used golangs channels as a third reversal function
*/
package main

import "fmt"

func ReverseFor(list []int) []int {
	for x, y := 0, len(list)-1; x < y; x, y = x+1, y-1 {
		list[x], list[y] = list[y], list[x]
	}
	return list
}

func ReverseRecursion(list []int) []int {
	if len(list) == 0 {
		return list
	}
	return append(ReverseRecursion(list[1:]), list[0])
}

func ReverseChannels(list []int, ch chan []int) {
	for x, y := 0, len(list)-1; x < y; x, y = x+1, y-1 {
		list[x], list[y] = list[y], list[x]
	}
	ch <- list
}

func main() {
	ch := make(chan []int)
	list := []int{
		1, 2, 3, 4, 5,
		6, 33, 8, 9, 10,
		11, 12, 13, 14,
	}
	fmt.Print(ReverseFor(list))
	fmt.Print(ReverseRecursion(list))
	go ReverseChannels(list, ch)
	reversed := <-ch
	fmt.Print(reversed)
}
