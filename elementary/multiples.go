package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			{
				fmt.Println(i)
			}
		case i%3 == 0:
			{
				fmt.Println(i)
			}
		case i%5 == 0:
			{
				fmt.Println(i)
			}
		}
	}
}
