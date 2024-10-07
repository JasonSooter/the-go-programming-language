package main

import (
	"bufio"
	"fmt"
	"os"
)

// `dup1` prints the text of each line that appears more than
// once in the standard input, preceded by its count.
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// cat ./ch1/dup1/OutstandingTasks.md | go run ./ch1/dup1/dup1.go
// Output: `2       ---`
