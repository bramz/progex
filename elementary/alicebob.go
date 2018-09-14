package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	buffer := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	name, err := buffer.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	name = strings.TrimSpace(name)
	switch {
	case name == "Alice":
		fmt.Println(name)
	case name == "Bob":
		fmt.Println(name)
	default:
		// do nothing
	}
}
