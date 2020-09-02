package main

import (
	"log"
	"os"
)

func main() {
	file, _ := os.Create("file_logger.out")
	defer file.Close()

	logger := log.New(file, "log example ", log.LstdFlags|log.Lshortfile)

	logger.Printf("This is log %s", "hello\n")
	logger.Println("This is another log")
}
