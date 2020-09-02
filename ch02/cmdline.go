package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "Alex", "Tell me your name")

var useSpanish bool

func init() {
	flag.BoolVar(&useSpanish, "spanish", false, "Use Spanish")
	flag.BoolVar(&useSpanish, "s", false, "Use Spanish")
}

func main() {
	flag.Parse()

	if useSpanish {
		fmt.Printf("Hola %s \n", *name)
	} else {
		fmt.Printf("Hello %s \n", *name)
	}

}
