package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("Hello, World!")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
