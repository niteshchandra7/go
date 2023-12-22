package main

import "fmt"

func main() {
	price, inStock := 100, true

	if price > 80 { // if condition only takes boolean expression
		fmt.Println("Too Expensive!")
	}
	_ = inStock
}
