package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dem4gus/roller"
)

func main() {
	if len(os.Args) == 1 {
		for {
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			result, err := roller.Roll(strings.TrimSpace(input))
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s", err)
				os.Exit(1)
			}
			fmt.Println(result)
		}
	}

	inputs := os.Args[1:]
	for _, in := range inputs {
		result, err := roller.Roll(in)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}
		fmt.Println(result)
	}
}
