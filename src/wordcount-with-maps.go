package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// []string
func WordCount(s string) map[string]int {
	//  method 1
	// fields := make([]string, len(strings.Fields(s)))
	// fields = strings.Fields(s)
	// range over fields
	// method 2
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		// fmt.Println(v)
		m[v] += 1
		// fmt.Println(m)
	}
	return m
	// return map[string]int{"x": 1}
}

func main() {
	// fmt.Println(WordCount("hello asd asd asd"))
	wc.Test(WordCount)
}
