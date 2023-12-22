package main

import "fmt"

type material uint

const (
	METAL material = iota
	PAPER
	GLASS
)

type truck struct {
	position uint
	contains material
	time     uint
}

func moveTruckTo(truck *truck, idx int, arr []int) {
	truck.time += uint(arr[idx] - arr[truck.position])
	truck.position = uint(idx)
}

func getMinTime(metalTruck *truck, paperTruck *truck, glassTruck *truck, garbage []string, arr []int) int {
	for i, v := range garbage {
		for _, c := range v {
			if c == 'G' {
				if glassTruck.position != uint(i) {
					moveTruckTo(glassTruck, i, arr)
				}
				glassTruck.time++
			} else if c == 'P' {
				if paperTruck.position != uint(i) {
					moveTruckTo(paperTruck, i, arr)
				}
				paperTruck.time++
			} else {
				if metalTruck.position != uint(i) {
					moveTruckTo(metalTruck, i, arr)
				}
				metalTruck.time++
			}
		}
	}
	return int(glassTruck.time) + int(metalTruck.time) + int(paperTruck.time)
}

func garbageCollection(garbage []string, travel []int) int {
	metalTruck := truck{position: 0, contains: METAL, time: 0}
	paperTruck := truck{position: 0, contains: PAPER, time: 0}
	glassTruck := truck{position: 0, contains: PAPER, time: 0}
	arr := getPrefixSumArr(travel)
	return getMinTime(&metalTruck, &paperTruck, &glassTruck, garbage, arr)
}

func getPrefixSumArr(travel []int) []int {
	ans := make([]int, len(travel)+1)
	ans[0] = 0
	for idx, v := range travel {
		ans[idx+1] = ans[idx] + v
	}
	return ans
}

func main() {
	garbage := []string{"G", "P", "GP", "GG"}
	travel := []int{2, 4, 3}
	fmt.Println(garbageCollection(garbage, travel))
}
