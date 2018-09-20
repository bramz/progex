/*
Write a program that prints a multiplication table for numbers up to 12.
*/
package main

import (
	"fmt"
)

const n = 12

func main() {
	for i := 1; i <= n; i++ {
		for s := 1; s <= 12; s++ {
			table := i * s
			fmt.Print(table)
			fmt.Printf(" | ")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
