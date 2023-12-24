package main

import "fmt"

func isPathCrossing(path string) bool {
	visited := make(map[[2]int]interface{})
	visited[[2]int{0, 0}] = nil
	currentLocation := [2]int{0, 0}
	directionMap := map[byte][2]int{
		'N': {0, 1}, 'E': {1, 0}, 'W': {-1, 0}, 'S': {0, -1},
	}
	for i := 0; i < len(path); i++ {
		nextLocation := currentLocation
		nextLocation[0] += directionMap[path[i]][0]
		nextLocation[1] += directionMap[path[i]][1]
		_, ok := visited[nextLocation]
		if ok {
			return true
		}
		visited[nextLocation] = nil
		currentLocation = nextLocation
	}
	return false
}

func main() {
	fmt.Println(isPathCrossing("NES"))
}
