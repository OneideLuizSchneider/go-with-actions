package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/joho/godotenv"
)

type Hello struct {
	Hello string `json:"hello"`
}

var helloWorld Hello = Hello{
	Hello: "world",
}

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("Go with cache :D")
	fmt.Printf("PR :D")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(helloWorld)
	})

	fmt.Println("Sever started...")

	http.ListenAndServe(":1337", nil)
}
