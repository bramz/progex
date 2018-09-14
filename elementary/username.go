package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buffer := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := buffer.ReadString('\n')
	fmt.Println(name)
}
