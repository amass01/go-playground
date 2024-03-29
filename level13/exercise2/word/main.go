package word

import "strings"

// UseCount is usage function
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count return the number of words in a quote
func Count(s string) int {
	return len(strings.Fields(s))
}
