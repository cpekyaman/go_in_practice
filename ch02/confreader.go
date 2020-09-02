package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Conf struct {
	Enabled bool
	Path    string
}

func main() {
	file, err := os.Open("conf.json")
	defer file.Close()
	if err != nil {
		log.Fatal("Could not load conf")
	}

	conf := Conf{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal("Could not parse conf")
	}
	fmt.Println(conf.Path)
}
