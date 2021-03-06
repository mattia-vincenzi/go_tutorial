package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	for _, elem := range strings.Fields(s) {
		wc[elem] += 1
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
