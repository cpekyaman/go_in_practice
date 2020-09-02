package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool, 1)
	for i := 1; i < 7; i++ {
		go worker(i, lock)
	}
	time.Sleep(5 * time.Second)
}

func worker(num int, lock chan bool) {
	fmt.Printf("Worker %d attempt to lock\n", num)
	lock <- true
	fmt.Printf("Worker %d acquired the lock\n", num)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Worker %d releasing the lock\n", num)
	<-lock
}
