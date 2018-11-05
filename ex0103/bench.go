package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(naive(os.Args[1:]))
}

func naive(args []string) string {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func join(args []string) string {
	return strings.Join(args, " ")
}
