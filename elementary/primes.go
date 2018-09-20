package main /*
Write a program that prints all prime numbers.
(Note: if your programming language does not support arbitrary size numbers,
printing all primes up to the largest number you can easily represent is fine too.)
*/
import (
	"fmt"
	"math"
)

func AllPrimes(number int) bool {
	for i := 2; i <= int(math.Floor(float64(number)/2)); i++ {
		if number%i == 0 {
			return false
		}
	}
	return number > 1
}

func main() {
	for i := 1; i <= 10000000000; i++ {
		if AllPrimes(i) {
			fmt.Printf("%v ", i)
		}
	}
}
