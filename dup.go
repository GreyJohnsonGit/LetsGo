package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// Definately confusing if someone actually used this
	for input.Scan() && input.Text() != "^C" {
		counts[input.Text()]++
	}
	// Ignoring errors
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
