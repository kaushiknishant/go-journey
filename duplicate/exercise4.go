package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fmt.Printf("%v\n", files)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		counts[input.Text()]++
	}
	
	if err := input.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

}
