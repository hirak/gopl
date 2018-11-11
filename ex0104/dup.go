package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		for line, _ := range counts {
			fileNames[line] = append(fileNames[line], "STDIN")
		}
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			defer f.Close()
			c := make(map[string]int)
			err = countLines(f, c)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			for line, n := range c {
				counts[line] += n
				fileNames[line] = append(fileNames[line], f.Name())
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(fileNames[line], ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]int) error {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	return input.Err()
}
