package main

import "fmt"

const (
	secondsInHour = 3600
)

func main() {
	fmt.Println(secondsInHour)
	var age int
	fmt.Println(age)
	s := "Learning GoLang"
	fmt.Printf("type of s is %T\n", s)

	// multiple declaration
	car, cost := "Audi", 5000
	fmt.Println(car, cost)

	var (
		salary    float64
		firstname string
		gender    bool
	)
	fmt.Println(salary, firstname, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	// multiple assignemnts
	var i, j int
	i, j = 5, 8
	j, i = i, j
	fmt.Println(i, j)
	fmt.Println(0.0 == 0.00)

}
