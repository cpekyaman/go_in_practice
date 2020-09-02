package main

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	baz()
}

func baz() {
	sb := make([]byte, 1024)
	runtime.Stack(sb, false)
	fmt.Printf("Stack trace \n %s \n", sb)
}
