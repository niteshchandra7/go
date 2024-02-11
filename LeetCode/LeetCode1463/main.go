package main

func cherryPickup(grid [][]int) int {
	dp := make([][][][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][][]int, len(grid[0]))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([][]int, len(grid))
			for k := 0; k < len(dp[i][j]); k++ {
				dp[i][j][k] = make([]int, len(grid[0]))
				for l := 0; l < len(dp[i][j][k]); l++ {
					dp[i][j][k][l] = -1
				}
			}
		}
	}
	return getCount(grid, 0, 0, 0, len(grid[0])-1, dp)
}

func getCount(grid [][]int, i int, j int, x int, y int, dp [][][][]int) int {

	if i == len(grid) || x == len(grid) || j < 0 || j == len(grid[0]) || y < 0 || y == len(grid[0]) {
		return 0
	}
	if dp[i][j][x][y] != -1 {
		return dp[i][j][x][y]
	}
	dir := []int{-1, 0, 1}
	maxSum := 0
	for k := 0; k < len(dir); k++ {
		temp1 := grid[i][j]
		temp2 := grid[x][y]
		if i == x && j == y {
			temp2 = 0
		}
		grid[i][j] = 0
		grid[x][y] = 0
		for n := 0; n < len(dir); n++ {
			maxSum = max(getCount(grid, i+1, j+dir[k], x+1, y+dir[n], dp)+temp1+temp2, maxSum)
		}
		grid[i][j] = temp1
		if i == x && j == y {
			grid[x][y] = temp1
		} else {
			grid[x][y] = temp2
		}
	}
	dp[i][j][x][y] = maxSum
	return dp[i][j][x][y]
}

func main() {

}
