package main

import (
	"fmt"
)

func main() {
	fmt.Printf("This is a Hello World Test!\n")
	fmt.Printf("%s\n", hello())
}

// Returns the appropriate greeting.
func hello() string {
	return "Hello, World!"
}
