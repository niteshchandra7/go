package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	i := 0
	j := 0
	count := 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
			j++
			count++
		} else {
			j++
		}
	}
	return count
}

func main() {

}
