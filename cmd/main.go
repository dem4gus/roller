package main

import (
	"fmt"
	"os"

	"github.com/dem4gus/roller"
)

func main() {
	result, err := roller.Roll(2, 6, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Rolling 2d6+2: %d", result)
}
