package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	words := strings.Split(s, " ")

	for _, word := range words {
		m[word] += 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
