package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)

}

func main() {
	var waitGroup sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"zeta",
		"pi",
		"epsilon",
		"theta",
	}
	waitGroup.Add(7)
	for i, x := range words {
		go printSomething(fmt.Sprintf("%d:%s", i, x), &waitGroup)
	}

	waitGroup.Wait()
	/*
		can put the go routine to sleep but is the worst way to handle concurrency
		time.Sleep(1 * time.Second)
	*/
	waitGroup.Add(1)

	/*
		A WaitGroup in Go is a simple way to wait for multiple goroutines to finish their execution.
		 It’s like a counter: when you launch a goroutine, you increase the counter, and when a goroutine finishes,
		 it decreases the counter. The Wait() function is used to block the execution of the program until the counter is zero,
		 meaning all the goroutines have finished.
	*/
	printSomething("second to be printed.", &waitGroup)
}
