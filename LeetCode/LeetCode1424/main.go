package main

func getSum(i int, j int, nums [][]int, c int) []int {
	var r int = len(nums)
	var ans []int
	for i < r && i >= 0 && j >= 0 && j < c {
		if j >= len(nums[i]) {
			i--
			j++
			continue
		}
		ans = append(ans, nums[i][j])
		i--
		j++
	}
	return ans
}

func findDiagonalOrder(nums [][]int) []int {
	var r int = len(nums)
	var c int = len(nums[r-1])
	var ans []int
	for i := 0; i < r-1; i++ {
		if len(nums[i]) > c {
			c = len(nums[i])
		}
	}
	for i := 0; i < r-1; i++ {
		ans = append(ans, getSum(i, 0, nums, c)...)
	}
	for j := 0; j < c; j++ {
		ans = append(ans, getSum(r-1, j, nums, c)...)
	}
	return ans
}
func main() {

}
