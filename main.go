package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: disk-usage-analyzer <path>")
		os.Exit(1) // exit code non-zero
	}

	path := os.Args[1]
	fmt.Println("Path:", path)
}
