package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args stores slice of string
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	var result, _ = strconv.ParseFloat(os.Args[0], 64)
	fmt.Printf("%T\n", result)
	

}
