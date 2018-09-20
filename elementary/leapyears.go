/*
Write a program that prints the next 20 leap years.
*/
package main

import (
	"fmt"
)

func LeapYear(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}

func main() {
	for i := 2018; i < 2100; i++ {
		if LeapYear(i) {
			fmt.Print(i)
			fmt.Printf("\n")
		}
	}
}
