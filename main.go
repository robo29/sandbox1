package main

import (
	"fmt"
)

const iteration = 5

func main() {
	fmt.Printf("This is a Hello World Test, interation %d.\n", iteration)
	fmt.Printf("%s\n", hello())
}

// Returns the appropriate greeting.
func hello() string {
	return "Hello, World!"
}
