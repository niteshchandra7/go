package main

func onesMinusZeros(grid [][]int) [][]int {
	r, c := len(grid), len(grid[0])
	onesRow := make([]int, r)
	onesCol := make([]int, c)
	ans := make([][]int, r)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			onesRow[i] += grid[i][j]
		}
	}
	for j := 0; j < c; j++ {
		for i := 0; i < r; i++ {
			onesCol[j] += grid[i][j]
		}
	}
	for i := 0; i < r; i++ {
		ans[i] = make([]int, c)
		for j := 0; j < c; j++ {
			ans[i][j] = 2*onesRow[i] - r + 2*onesCol[j] - c
		}
	}
	return ans
}

func main() {

}
