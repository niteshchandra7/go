package main

func getTotalOnes(s string) int {
	count := 0
	for _, v := range s {
		if v == '1' {
			count++
		}
	}
	return count
}

func maxScore(s string) int {
	ans := 0
	totalOnes := getTotalOnes(s)
	left_0 := 0
	left_1 := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '1' {
			left_1++
		} else {
			left_0++
		}
		score := left_0 + (totalOnes - left_1)
		if score > ans {
			ans = score
		}
	}
	return ans
}

func main() {

}
