package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func(c chan int) {
		//defer close(c)
		c <- 10
		c <- 20
		c <- 30
		time.Sleep(2 * time.Second)
		c <- 40
	}(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
