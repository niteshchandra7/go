package main

const (
	mod = 1e9 + 7
)

func numRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	return getNumRollsToTarget(n, k, target, dp)
}

func getNumRollsToTarget(n int, k int, target int, dp [][]int) int {
	if n == 0 && target == 0 {
		return 1
	} else if target <= 0 || n < 0 {
		return 0
	}
	if dp[n][target] != -1 {
		return dp[n][target]
	}
	count := 0
	for j := 1; j <= k; j++ {
		count = (count%mod + getNumRollsToTarget(n-1, k, target-j, dp)%mod) % mod
	}
	dp[n][target] = count
	return count
}

func main() {

}
