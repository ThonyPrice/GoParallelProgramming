/*
Task 2.3 - Slices, from https://tour.golang.org/moretypes/23
DD1396 - Parallel and concurrent programming @ KTH
Author: Thony Price
Last revision: 2017-03-22
*/

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount takes a string and return a map with each unique word
// as keys and the occurence of each as values
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	wordList := strings.Fields(s)
	for _, v := range wordList {
		m[v] = count(v, wordList)
	}
	return m
}

// count takes a string and a list of strings and return and integer of
// how many times the string occures in the list
func count(s string, l []string) int {
	x := 0
	for _, v := range l {
		if s == v {
			x++
		}
	}
	return x
}

func main() {
	wc.Test(WordCount)
}
