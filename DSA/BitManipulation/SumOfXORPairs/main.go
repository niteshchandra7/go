package main

import "fmt"

// Given an array of N elements. Calculate the sum of XOR of all possible pairs.

func getSumOfXORPairs(arr []int) int {
	bitCountArr := make([][]int, 32)
	for i := 0; i < 32; i++ {
		count0 := 0
		count1 := 0
		for j := 0; j < len(arr); j++ {
			if (arr[j] & (1 << i)) != 0 {
				count1++
			} else {
				count0++
			}
		}
		bitCountArr[i] = []int{count0, count1}
	}
	sum := 0
	for k, v := range bitCountArr {
		sum += ((1 << k) * (v[0] * v[1]))
	}
	return sum
}

func main() {

	arr := []int{3, 2, 8, 5, 6}
	fmt.Println(getSumOfXORPairs(arr) * 2)

}
