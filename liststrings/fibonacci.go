package main /*
Write a function that computes the list of the first 100 Fibonacci numbers.
The first two Fibonacci numbers are 1 and 1.
The n+1-st Fibonacci number can be computed by adding the n-th and the n-1-th Fibonacci number.
The first few are therefore 1, 1, 1+1=2, 1+2=3, 2+3=5, 3+5=8.
*/import "fmt"

func Fibonacci(n int) int {
	fn := make(map[int]int)
	var f int
	for i := 0; i <= n; i++ {
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	fmt.Print(Fibonacci(100))
}
