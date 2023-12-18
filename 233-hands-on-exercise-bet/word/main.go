// Package word provides custom functions with words in a string.
package word

import "strings"

// UseCount counts quantity of repeating words in the quote
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count counts the words in the quote
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
