package main

import "fmt"

func main() {

	var cities []string
	fmt.Println(len(cities))   // 0
	fmt.Println(cap(cities))   // 0
	fmt.Println(cities == nil) // true

	cities = append(cities, "Hello")
	places := cities
	fmt.Printf("%p\n", &cities)
	fmt.Printf("%p\n", &places)
	places[0] = "Bye"
	fmt.Printf("%#v\n", cities)
	fmt.Printf("%#v\n", places)

	//arr1 := []int{1, 2}
	//arr2 := []int{1, 2}
	//fmt.Println(arr1 == arr2) // slices con't be compared

}
