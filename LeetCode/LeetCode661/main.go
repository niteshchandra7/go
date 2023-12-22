package main

func imageSmoother(img [][]int) [][]int {
	r, c := len(img), len(img[0])
	ans := make([][]int, len(img))
	for i := 0; i < r; i++ {
		ans[i] = make([]int, c)
		for j := 0; j < c; j++ {
			ans[i][j] = getAverage(img, i, j)
		}
	}
	return ans
}

func getAverage(img [][]int, i int, j int) int {
	r, c := len(img), len(img[0])
	arr := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 0},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	count := 0
	total := 0
	for k := 0; k < len(arr); k++ {
		x := i + arr[k][0]
		y := j + arr[k][1]
		if x < 0 || y < 0 || x >= r || y >= c {
			continue
		}
		count++
		total += img[x][y]
	}
	return total / count
}

func main() {

}
