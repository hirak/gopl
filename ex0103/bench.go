package main

import (
	"strings"
)

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
