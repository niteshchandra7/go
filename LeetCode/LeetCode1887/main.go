package main

import (
	"fmt"
	"sort"
)

func reductionOperations(nums []int) int {
	ans, weight := 0, 0
	sort.Slice(nums, func(a, b int) bool {
		return (nums[a] < nums[b])
	})
	numsMap := make(map[int]int)
	for _, v := range nums {
		_, ok := numsMap[v]
		if !ok {
			numsMap[v] = 1
		} else {
			numsMap[v] += 1
		}
	}
	for i := 0; i < len(nums); {
		j := i
		for j < len(nums) && nums[j] == nums[i] {
			j++
		}
		ans += (weight * numsMap[nums[i]])
		weight++
		i = j
	}
	return ans

}

func main() {
	var size int
	fmt.Printf("Enter size of the array: ")
	fmt.Scanf("%d", &size)
	var nums []int = make([]int, size)
	for idx, _ := range nums {
		fmt.Printf("Enter %d index value: ", idx)
		fmt.Scanf("%d", &nums[idx])

	}
	fmt.Println(reductionOperations(nums))

}
