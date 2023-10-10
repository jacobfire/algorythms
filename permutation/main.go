package main

import (
	"fmt"
	"strings"
)

func main() {
	permutations := AllPermutations("abc")
	fmt.Println(permutations)
}

func AllPermutations(s string) map[string]string {
	p := make(map[string]string)
	for i := 0; i < len(s); i++ {
		b, a, _ := strings.Cut(s, string(s[i]))
		for _, v := range exclude(string(s[i]), a+b) {
			if _, ok := p[v]; !ok {
				p[v] = v
			}
		}
	}

	return p
}

func exclude(symbol, left string) []string {
	var options []string
	for i := 0; i <= len(left); i++ {
		option := left[:i] + symbol + left[i:]
		options = append(options, option)
	}

	return options
}
