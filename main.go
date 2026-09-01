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

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading directory:", err)
		os.Exit(1) // exit code non-zero
	}

	var total int64

	for _, entry := range entries {
		if entry.IsDir() {
			continue // skip directories
		}

		info, err := entry.Info()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error getting file info", entry.Name(), ":", err)
			continue // skip this entry

		}
		total += info.Size()
	}
	fmt.Println("Total size:", total)
}
