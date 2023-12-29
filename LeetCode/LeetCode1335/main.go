package main

import (
	"fmt"
	"math"
)

func solve(arr []int, idx int, d int, dp [][]int) int {
	max := math.MinInt32
	ans := math.MaxInt32
	if idx == len(arr) {
		if d == 0 {
			return 0
		} else {
			return ans
		}
	}
	if d == 0 {
		return ans
	}
	if dp[idx][d] != math.MaxInt32 {
		return dp[idx][d]
	}
	for i := idx; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		temp := max + solve(arr, i+1, d-1, dp)
		if temp < ans {
			ans = temp
		}
	}
	dp[idx][d] = ans
	return ans
}

func minDifficulty(jobDifficulty []int, d int) int {
	dp := make([][]int, len(jobDifficulty))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, d+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	ans:=solve(jobDifficulty, 0, d, dp)
	if ans==math.MaxInt32{
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minDifficulty([]int{6, 5, 4, 3, 2, 1}, 2))
}
