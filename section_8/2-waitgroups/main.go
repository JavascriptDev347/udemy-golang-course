package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(text string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("Say Hello: ", text)
}

func main() {
	/*
	 1. Add outside of the goroutine
	 2. You must decrease the counter by calling wg.Done inside the goroutine
	 3. Do not forget to call wg.Wait()
	 4. Always pass a reference/pointer of the wait group variable instead of a copy
	*/
	var wg sync.WaitGroup
	//wg.Add(3)
	//fmt.Println("Hello from main goroutine")
	//
	//go sayHello("Go Routine 1", time.Second, &wg)
	//go sayHello("Go Routine 2", 2*time.Second, &wg)
	//go sayHello("Go Routine 3", 5*time.Second, &wg)

	totalJobs := 5

	for i := 0; i < totalJobs; i++ {
		wg.Add(1)
		go sayHello(fmt.Sprintf("JOB %d", i), time.Second, &wg)
	}
	fmt.Println("Last message from main goroutine")

	wg.Wait()
}
