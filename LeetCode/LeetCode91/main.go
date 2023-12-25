package main

import "strconv"

func numDecodings(s string) int {
	dp := make([]int, len(s))
	for i := range dp {
		dp[i] = -1
	}
	return getDecodingsCount([]rune(s), 0, dp)
}

func getDecodingsCount(arr []rune, i int, dp []int) int {
	if i >= len(arr) {
		return 1
	} else if rune(arr[i]) == '0' {
		return 0
	} else if dp[i] != -1 {
		return dp[i]
	}
	count := getDecodingsCount(arr, i+1, dp)
	if i+1 < len(arr) {
		str := string(arr[i]) + string(arr[i+1])
		num, _ := strconv.Atoi(str)
		if num <= 26 {
			count += getDecodingsCount(arr, i+2, dp)
		}
	}
	dp[i] = count
	return dp[i]
}

func main() {

}
