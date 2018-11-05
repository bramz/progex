/*
Write a function that returns the elements on odd positions in a list.
*/
package main

import "fmt"

func ReturnOddPositions(list []int) {
	for i := 0; len(list) > i; i++ {
		if i%2 != 0 {
			newlist := []int{
				list[i],
			}
			i++
			fmt.Print(newlist)
		}
	}
}

func main() {
	list := []int{
		1, 2, 3, 4, 5,
		6, 29, 83, 2,
		111, 1221, 312, 313,
	}
	ReturnOddPositions(list)

	/*
		for i := 0; len(list) > i; i++ {
			if i%2 != 0 {
				newlist := []int{
					list[i],
				}
				i++
				fmt.Print(newlist)
			}
		}
	*/
}
