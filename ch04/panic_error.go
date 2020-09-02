package main

import (
	"errors"
	"fmt"
)

var DivByZero = errors.New("Divide By Zero")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("We got panic %s\n", err)
		}
	}()

	_, err := safe_div(1, 0)
	if err != nil {
		fmt.Println("safe_div err", err)
	}

	res := div(2, 0)
	fmt.Println("2 / 0 is", res)
}

func safe_div(num int, den int) (int, error) {
	if den == 0 {
		return 0, DivByZero
	}
	return div(num, den), nil
}

func div(num int, den int) int {
	return num / den
}
