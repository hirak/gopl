package main

import "testing"

func TestFuncs(t *testing.T) {
	testcases := map[string][]string{
		"aaa bbb ccc": []string{"aaa", "bbb", "ccc"},
		"aaa bbb":     []string{"aaa", "bbb"},
		"aaa":         []string{"aaa"},
		"":            []string{},
	}
	funcs := map[string]func([]string) string{
		"naive": naive,
		"join":  join,
	}
	for expect, input := range testcases {
		for funcName, f := range funcs {
			actual := f(input)
			if actual != expect {
				t.Errorf("%s: want: %s, got: %s", funcName, expect, actual)
			}
		}
	}
}
