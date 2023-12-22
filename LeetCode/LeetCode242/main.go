package main

import (
	"sort"
	"strings"
)

func main() {

}

func isAnagram(s string, t string) bool {
	s1 := []byte(s)
	t1 := []byte(t)
	sort.Slice(s1, func(i int, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(t1, func(i int, j int) bool {
		return t1[i] < t1[j]
	})
	s2 := string(s1)
	t2 := string(t1)
	return strings.EqualFold(s2,t2)

}
