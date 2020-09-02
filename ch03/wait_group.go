package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for num := 1; num < 5; num++ {
		wg.Add(1)
		go task(&wg, num)
	}
	fmt.Println("Waiting for tasks")
	wg.Wait()
	fmt.Println("Done")
}

func task(wg *sync.WaitGroup, num int) {
	fmt.Printf("Task %d runnıng\n", num)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Task %d finished\n", num)
	wg.Done()
}
