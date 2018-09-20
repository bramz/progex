/*
Write a guessing game where the user has to guess a secret number.
After every guess the program tells the user whether their number was too large or too small.
At the end the number of tries needed should be printed.
It counts only as one try if they input the same number multiple times consecutively.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const limit = 20

func rando(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	var guess int
	random := rando(1, limit)
	attempts := 0

	for guess != random {
		fmt.Printf("Guess a number between 1..%d: ", limit)
		fmt.Scanf("%d", &guess)
		attempts++

		switch {
		case guess > random:
			{
				fmt.Print("Your number is too high!\n")
			}
		case guess < random:
			{
				fmt.Print("Your number is too low!\n")
			}
		case guess == random:
			{
				fmt.Printf("Your answer was correct in %v attempts.", attempts)
			}
		}
	}
}
