package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	
	for _, sub := range strings.Fields(s)  {
		
		if ct, ok := m[sub]; ok {
			m[sub] = ct + 1
		} else {
			m[sub] = 1
		}
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}
