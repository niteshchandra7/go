package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

type Balloon struct {
	Color byte
	Time  int
}

func minCost(colors string, neededTime []int) int {
	balloons := make([]Balloon, len(colors))
	ans := 0
	for i := 0; i < len(balloons); i++ {
		balloons[i] = Balloon{
			Color: colors[i],
			Time:  neededTime[i],
		}
	}
	s := stack.New()
	s.Push(balloons[0])
	for i := 1; i < len(balloons); i++ {
		top := s.Peek().(Balloon)
		if s.Len() != 0 && top.Color == balloons[i].Color {
			if top.Time < balloons[i].Time {
				s.Pop()
				ans += top.Time
				s.Push(balloons[i])
			} else {
				ans += balloons[i].Time
			}
		} else {
			s.Push(balloons[i])
		}

	}
	return ans
}

func main() {
	fmt.Println(minCost("aabaa", []int{1, 2, 3, 4, 1}))
}
