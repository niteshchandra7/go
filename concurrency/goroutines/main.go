package main

import (
	"fmt"
	"runtime"
	"sync"
)

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f1(goroutine) exceution started")

	for i := 0; i < 3; i++ {
		fmt.Println("f1,i=", i)
	}
	fmt.Println("f1 executuon ended")
}

func f2() {
	fmt.Println("f2(goroutine) exceution started")

	for i := 0; i < 3; i++ {
		fmt.Println("f2,i=", i)
	}
	fmt.Println("f2 executuon ended")
}

func main() {
	fmt.Println("main execution started")
	fmt.Println("No of CPUs:", runtime.NumCPU())
	fmt.Println("No fo go routines:", runtime.NumGoroutine())

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // num of max os threads accessible to your go program
	// by default it is equal to number of CPUs
	var wg sync.WaitGroup
	wg.Add(1)
	go f1(&wg)
	fmt.Println("No fo go routines after go f1():", runtime.NumGoroutine())
	f2()
	//time.Sleep(2 * time.Second)
	wg.Wait()
	fmt.Println("main executiion stopped")

}
