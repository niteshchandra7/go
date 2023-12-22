package main

import "fmt"

func main() {
	const days int = 7 // constant works even when not used
	const pi float64 = 3.14

	// constant belongs to compile time while variable belongs to run time

	const n, m int = 4, 5 // typed constant
	const n1, m1 = 6, 7   // untyped constant
	//a:=2.2*n // error
	a := 2.2 * n1 // ok
	fmt.Println(a)
	const (
		min1 = -500
		min2
		min3
	) // group constants

	fmt.Println(min1, min2, min3)

	// Constants rules
	//1. you can not change a constant
	const temp int = 100
	// temp = 50 //Error

	//2. You can not initialize a constant at run time
	// const power = math.Pow(2,3)

	//3. You can not use a varialbe to intialise a constant
	//t := 5
	//const tc = t

	//4.
	const l1 = len("hello") //len is compile time function

	var arr [5]int = [5]int{0, 1, 2, 3, 4}
	fmt.Printf("%T, %v\n", arr, arr)

	// defined types
	type age int // int is its underlying types
	type oldAge age // int is its underlying types
	type veryOldAge age // int is its underlying types

	// alias, new name same type
	type second = uint

}
