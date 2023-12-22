package main

import "sort"

func check(nums []int, s int, e int) bool {
	ans := make([]int, e-s+1)
	j := 0
	for i := s; i <= e; i++ {
		ans[j] = nums[i]
		j++
	}
	sort.Slice(ans, func(a, b int) bool {
		return ans[a] < ans[b]
	})
	return isArithmetic(ans)
}

func isArithmetic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	d := nums[1] - nums[0]
	for i := 1; i < len(nums); i++ {
		if (nums[i] - nums[i-1]) != d {
			return false
		}
	}
	return true
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		ans[i] = check(nums, l[i], r[i])
	}
	return ans
}

func main() {

}
