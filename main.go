package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	//"github.com/joho/godotenv"
	//"github.com/joho/godotenv/cmd/godotenv"
	_ "github.com/joho/godotenv/autoload"
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
	fmt.Printf("More logs...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(helloWorld)
	})

	fmt.Println("Sever started...")

	http.ListenAndServe(":1337", nil)
}
