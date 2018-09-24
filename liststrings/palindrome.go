/*
Write a function that tests whether a string is a palindrome.
*/
package main

import "fmt"

func isPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	word := "racecar"
	if isPalindrome(word) == true {
		fmt.Print("Is palindrome")
	} else {
		fmt.Print("Is not palindrome")
	}
}
