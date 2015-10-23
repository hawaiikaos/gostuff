package main

import (
	"golang.org/x/tour/wc"
	//"fmt"
)

func WordCount(s string) map[string]int {
	
	space := rune(32)
	r := []rune(s)
	i := 0
	j := 0
	var p []string
	p = append(p,"")
	for i < len(r) {
		if r[i] == space {
			j++
			p = append(p,"")
		} else {
			p[j] += string(r[i])
		}
		i++
	}
	m := make(map[string]int)
	m[p[0]] = 1
	
	i = 0
	for i < len(p) {
		wordcount := 0
		k := 0
		for k < len(p) {
			if p[i] == p[k] {
				wordcount++
			}
			k++
		}
		m[p[i]] = wordcount
		i++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
