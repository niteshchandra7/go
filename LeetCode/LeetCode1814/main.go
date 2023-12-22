package main

import (
	"fmt"
	"strconv"
)

const (
	MOD = 1e9 + 7
)

func getRev(n int) int {
	str := strconv.Itoa(n)
	var revStr string
	idx := len(str) - 1
	for idx >= 0 && str[idx] == '0' {
		idx--
	}
	if idx < 0 {
		return 0
	}
	for ; idx >= 0; idx-- {
		revStr += string(str[idx])
	}
	ans, _ := strconv.Atoi(revStr)
	return ans

}

func countNicePairs(nums []int) int {
	comp := make([]int, len(nums))
	var ans int = 0
	mp := make(map[int]int)
	for i, v := range nums {
		comp[i] = v - getRev(v)
		count, ok := mp[comp[i]]
		if ok {
			ans = ((ans % MOD) + (count % MOD)) % MOD
		}
		mp[comp[i]]++
	}
	return ans
}

func main() {
	fmt.Println(countNicePairs([]int{42, 11, 1, 97}))
}
