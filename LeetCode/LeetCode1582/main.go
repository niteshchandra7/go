package main

func main() {

}

func numSpecial(mat [][]int) int {
	ans := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 && check(mat, i, j) {
				ans++
			}
		}
	}
	return ans
}

func check(mat [][]int, x int, y int) bool {
	count := 0
	for j := 0; j < len(mat[0]); j++ {
		if mat[x][j] == 1 {
			count++
		}
	}
	if count != 1 {
		return false
	}
	count = 0
	for i := 0; i < len(mat); i++ {
		if mat[i][y] == 1 {
			count++
		}
	}
	return count == 1
}
