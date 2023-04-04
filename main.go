package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var htmlStr string

func main() {
	fmt.Println("start")

	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	htmlStr = string(data)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, htmlStr)
}
