package main /*
Write a function that takes a number and returns a list of its digits. So for 2342 it should return [2,3,4,2].
*/import "fmt"

func ListDigits(n int) []int {
	var list []int
	for i := 1; i <= n; i++ {
		list = []int{i}
	}
	return list
}

func main() {
	fmt.Print(ListDigits(12345))
}
