/*
Write a program that asks the user for a number n
and gives them the possibility to choose between
computing the sum and computing the product of 1,…,n.
*/
package main

import (
	"fmt"
)

func main() {
	var n, sum, product int
	var ps string

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	fmt.Print("Product or Sum? ")
	fmt.Scan(&ps)

	if ps == "sum" {
		for i := 1; i <= n; i++ {
			sum += i
		}
		fmt.Print(sum)
	} else {
		product = 1
		for i := 1; i <= n; i++ {
			product *= i
		}
		fmt.Print(product)
	}
}
