package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var builder strings.Builder
	for index, argument := range os.Args {
		fmt.Fprintf(&builder, "%d: %s\n", index, argument)
	}
	fmt.Println(builder.String())
}
