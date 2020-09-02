package main

import (
	"fmt"
	"time"
)

func count(c chan int) {
	for n := range c {
		fmt.Println("n out ", n)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	c := make(chan int)
	a := []int{3, 4, 6, 1, 1, 6, 8}

	go count(c)
	for _, n := range a {
		c <- n
	}
	fmt.Println("Waiting")
	time.Sleep(time.Millisecond * 5)
	close(c)
}
