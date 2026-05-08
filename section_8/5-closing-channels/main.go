package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			j, more := <-jobs
			if more {
				println("Received job", j)
			} else {
				println("Received all jobs")
				return
			}
		}
	}(&wg)
	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sending job", i)
	}

	close(jobs)

	wg.Wait()
}
